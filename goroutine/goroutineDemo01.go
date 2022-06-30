package main

import "fmt"

func Hello()  {
	fmt.Println("sdsadad")
}

func main() {
	go Hello()
	fmt.Println("goroutine")
}
