package main

import (
	"context"
	"fmt"
	"runtime/debug"
	"time"
)

func main() {
	// 希望捕获所有所有 panic
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	group := NewPanicGroup(1)
	// 启动新协程
	group.Go(func() {
		panic("有问题")
	})

	// 等待一下，不然协程可能来不及执行
	time.Sleep(1 * time.Second)
	err := group.Wait(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("这条消息打印不出来")
}

type PanicGroup struct {
	panics chan Panic // 协程 panic 通知信道
	dones  chan int   // 协程完成通知信道
	jobs   chan int   // 协程并发数量控制信道
	jobN   int32      // 协程并发数量
}

// Panic 子协程 panic 会被重新包装，添加调用栈信息
type Panic struct {
	R     interface{} // recover() 返回值
	Stack []byte      // 当时的调用栈
}

func (p Panic) String() string {
	return fmt.Sprintf("%v\n%s", p.R, p.Stack)
}

// NewPanicGroup 创建新的协程分组
func NewPanicGroup(maxConcurrent int) *PanicGroup {
	return &PanicGroup{
		panics: make(chan Panic, 8),
		dones:  make(chan int, 8),
		jobs:   make(chan int, maxConcurrent),
	}
}

// Go 新建协程并执行 f()，需要跟 Wait 在同一协程使用
//
// 不同的业务场景需要不同的初始参数。我们没法预测参数的数量和类型，索性定义
// f() 不接受任何参数，具体参数在使用的时候通过闭包捕获。
//
// 注意，所有协程的业务逻辑都需要透传跟 Wait() 调用相同的 ctx 并处理通知逻辑。
//
// 协程在执行的时候可能有两种报错 error 和 panic。
//
// 业务代码使用 Wait() 方法等等，因为可能有多个协程产生 error，我们没法确定
// error 的数量，也就不能很好定义 f() 的返回值，索性规定 f() 没有返回值，具体
// 的业务报错由程序员自行控制。
//
// 如果产生了 panic，这种错误一般不能恢复，Wait()方法直接将 panic 重新拋出。
// 业务代码可以决定是否处理对应的 panic 或者让框架来处理。
func (g *PanicGroup) Go(f func()) *PanicGroup {
	g.jobN++
	go func() {
		g.jobs <- 1
		defer func() {
			<-g.jobs
			// go 语言只能在自己的协程中捕获自己的 panic
			// 如果不处理，整个*进程*都会退出
			if r := recover(); r != nil {
				g.panics <- Panic{R: r, Stack: debug.Stack()}
				// 如果发生 panic 就不再通知 Wait() 已完成
				// 不然就可能出现 g.jobN 为 0 但 g.panics 非空
				// 的情况，此时 Wait() 方法需要在正常结束的分支
				// 中再额外检查是否发生了 panic，非常麻烦
				return
			}
			g.dones <- 1
		}()
		f()
	}()

	return g
}

// Wait 等待所有协程结束
// 保证协程内产生 panic 不会导致整个进程退出，但本方法依然会向上拋出对应 panic
// 如果 ctx 被取消，本方法会返回对应错误
// 如果没有任务等等会直接 panic!!!
func (g *PanicGroup) Wait(ctx context.Context) error {
	if g.jobN == 0 {
		panic("no job to wait")
	}

	for {
		select {
		case <-g.dones: // 协程正常结束
			g.jobN--
			if g.jobN == 0 {
				return nil
			}
		case p := <-g.panics: // 协程有 panic
			panic(p)
		case <-ctx.Done():
			// 整个 ctx 结束，超时或者调用方主动取消
			// 子协程应该共用该 ctx，都会收到相同的结束信号
			// 不需要在这里再去通知各协程结束（实现起来也麻烦）
			return ctx.Err()
		}
	}
}
