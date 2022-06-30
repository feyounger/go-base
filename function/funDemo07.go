package main

import "fmt"

type HandleFunc func(string)

func main() {
	//defer func() {
	//	if err := recover(); err!=nil{
	//		fmt.Println(err)
	//	}
	//}()
	//ch := make(chan int, 10)
	//close(ch)
	//ch <- 100
	handleFunc := HandleFunc(func(s string) {
		fmt.Println(s)
	})
	handleFunc("1")
}
