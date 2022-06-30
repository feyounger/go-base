package main

import "fmt"

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}
func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}
func main() {
	s := make([]int, 2, 4)
	s[0] = 1
	s[1] = 1
	fmt.Println("len", len(s))
	fmt.Println("cap", cap(s))
	fmt.Println(&s)
	newS := myAppend(s)
	fmt.Println(s)
	fmt.Println(newS)
	s = newS
	myAppendPtr(&s)
	fmt.Println(s)
}
