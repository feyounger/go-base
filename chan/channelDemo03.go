package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func setup() {
	// 这里面有一些初始化的操作
}

func main() {
	setup()

	// 用于监听服务退出, 这里使用了两个 goroutine，所以 cap 为2
	done := make(chan error, 2)

	// 无缓冲的通道，用于控制服务退出，传入同一个 stop，做到只要有一个服务退出了那么另外一个服务也会随之退出
	stop := make(chan struct{}, 0)

	// for debug
	go func() {
		//  pprof 传递一个 channel
		fmt.Println("pprof start...")
		done <- pprof(stop)
		fmt.Printf("err1:%v\n", done)

	}()

	// 主服务
	go func() {
		fmt.Println("app start...")
		done <- app(stop)
		fmt.Printf("err2:%v\n", done)
	}()

	// stopped 用于判断当前 stop 的状态
	var stopped bool

	// 这里循环读取 done 这个 channel
	// 只要有一个退出了，我们就关闭 stop channel
	for i := 0; i < cap(done); i++ {

		// 对于有缓冲的chan, chan中无值会一直处于阻塞状态
		// 对于app 服务会一直阻塞状态，不会有 数据写入到done 通道，只有在5s后，模拟的 pprof会有err写入chan,此时才会触发以下逻辑
		if err := <-done; err != nil {
			log.Printf("server exit err: %+v", err)
		}

		if !stopped {
			stopped = true
			// 通过关闭 无缓冲的channel 来通知所有的 读 stop相关的goroutine退出
			close(stop)
		}
	}
}

// http 服务
func app(stop <-chan struct{}) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	return server(mux, ":8080", stop)
}

func pprof(stop <-chan struct{}) error {
	// 注意这里主要是为了模拟服务意外退出，用于验证一个服务退出，其他服务同时退出的场景
	// 因为这里没有返回err， 所以done chan中无法接收到值， 主程序中会一直阻塞住
	go func() {
		server(http.DefaultServeMux, ":8081", stop)
	}()

	time.Sleep(5 * time.Second)
	// 模拟出错
	return fmt.Errorf("mock pprof exit")
}

// 启动一个服务
func server(handler http.Handler, addr string, stop <-chan struct{}) error {

	s := http.Server{
		Handler: handler,
		Addr:    addr,
	}

	// 这个 goroutine 控制退出，因为 stop channel 只要close或者是写入数据，这里就会退出
	go func() {
		// 无缓冲channel等待，写入或者关闭
		<-stop
		log.Printf("server will exiting, addr: %s", addr)
		// 此时 httpApi 服务就会优雅的退出
		s.Shutdown(context.Background())
	}()

	// 没有触发异常的话，会一直处于阻塞
	return s.ListenAndServe()
}