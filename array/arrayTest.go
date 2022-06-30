package main

import "fmt"

func main() {
	//arr := [...][3]int{{1,1,1},{2,2,2},{3,3,3}}
	//for k1, v1 := range arr {
	//	for k2, v2 := range v1 {
	//		fmt.Printf("(%d,%d)=%d ",k1,k2,v2)
	//	}
	//	fmt.Println()
	//}
	arr := [5]int{}
	printArr(&arr)
	fmt.Println(arr)
	fmt.Printf("arr: %p\n",&arr)
	//arr2 := [...]int{1,2,3,4,5}
	//printArr(&arr2)
	//fmt.Printf("arr2: %p\n",&arr2)
}

func printArr(arr *[5]int) {
	arr[0] = 1100
	for k1, v1 := range arr {
		fmt.Printf("(%d,%d)",k1,v1)
	}
	fmt.Printf("arr: %p\n",&arr)
	fmt.Println()
}