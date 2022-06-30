package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	// 定义 goroutine 数量
	goroutineCount = 4
	// 任务数
	taskCount = 10
	wg sync.WaitGroup
)

func main()  {
	// 有缓冲通道 类型为 string, 缓冲数 taskCount 为 10
	tasks := make(chan string, taskCount)
	wg.Add(goroutineCount)

	for gr:=1; gr<=goroutineCount; gr++{
		// 开启多个 goroutine 执行任务
		go worker(tasks, gr)
	}

	// 存放任务到任务通道中
	for task :=1; task <= taskCount; task++ {
		tasks <- fmt.Sprintf("Task %d\n", task)
	}
	// 执行完所有的 goroutine 后，关闭通道
	close(tasks)

	// 阻塞主线程，等到所有 goroutine 执行完毕在往下执行
	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for  {
		// 从任务通道中取任务 和 通道 是否关闭的状态
		task, ok := <-tasks
		// 通道已经关闭
		if !ok {
			fmt.Printf("Worker %d: shutdown\n", worker)
			return
		}
		// 模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		fmt.Printf("Worker %d started %s", worker, task)
		// 完成工作
		fmt.Printf("Worker %d completed %s\n", worker, task)
	}
}