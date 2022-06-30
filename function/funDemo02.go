package main

import "fmt"

func main() {
	slice1 := []int{1,2,3,4,5,6}
	println(argTest("sum: %d", slice1...))
	fmt.Println(slice1[0])
}

func argTest(s string, n ...int) string{
	var sum int
	n[0] = 10
	for _, v := range n {
		sum += v
	}
	return fmt.Sprint(s,sum)
}
