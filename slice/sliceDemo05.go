package main

import "fmt"

func main() {



	array := [4]int{10, 20, 30, 40}
	slice := array[0:3]
	foo1 := foo(slice)
	fmt.Println(array[0])
	fmt.Println(slice[0])
	fmt.Println(foo1[0])
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	fmt.Printf("After array = %v\n", array)
}

func foo(slice []int) []int {
	slice[0] = 100
	return slice
}
