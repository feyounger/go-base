package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := array[0:2]
	//slice2 := slice1[0:3]
	//slice2[1] = 222
	fmt.Println(array)
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))
	//fmt.Println(slice2)
}