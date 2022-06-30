package main

import (
	"fmt"
)

func main() {
	arr := [...]int{1,2,3,4,5,6,7,8,9}
	slice1 := arr[:4]
	fmt.Println(slice1)
	fmt.Println(cap(slice1))
	fmt.Println(len(slice1))

	slice2:= make([]int, 5, 10)
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5,6,7,8)
	fmt.Println("%v",slice2)
	fmt.Println(cap(slice2))
	fmt.Println(len(slice2))

	//myStr := []string{89: "a"}
	//fmt.Println(myStr)
	//fmt.Println(len(myStr))
	//fmt.Println(cap(myStr))
	//
	//var myNum  []int
	//fmt.Println(myNum)
	//fmt.Println(len(myNum))
	//fmt.Println(cap(myNum))
}
