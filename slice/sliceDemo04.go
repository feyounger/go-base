package main

import "fmt"

func main() {
	arrayA := []int{100, 200}
	testArrayPoint(&arrayA)   // 1.传数组指针
	arrayB := arrayA[:]
	fmt.Printf("arrayB : %p , %v\n", &arrayB, arrayB)
	testArrayPoint(&arrayB)   // 2.传切片
	fmt.Printf("arrayA : %p , %v\n", &arrayA, arrayA)
}

func testArrayPoint(arr *[]int) {
	fmt.Printf("func Array : %p , %v\n", arr, *arr)
	(*arr)[1] += 100
}
