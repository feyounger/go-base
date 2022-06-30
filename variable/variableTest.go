package main

import "fmt"

func init() {

}

func main() {
	x, _ := foo()
	_, y := foo()
	fmt.Println(x)
	fmt.Println(y)
}

func foo() (int, string) {
	return 10,"Feyoung"
}

