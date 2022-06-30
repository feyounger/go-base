package main

import (
	"fmt"
	"time"
)

func main(){
	tc:= make(chan int, 1)
	go func() {
		for i := 0; i <10; i++ {
			time.Sleep(time.Second)
			tc <- i
		}
	}()
	//<-tc
	fmt.Println("tick")
	select {

	}
	//tk:= time.Tick(time.Second)
	//<- tk
	//select {
	//
	//}
}
