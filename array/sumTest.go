package main

import "fmt"

func sumArr(arr *[5]int) {
	fmt.Printf("arr: %p\n",&arr)
	arr[0]=12132
}

func main() {
	arr := [5]int{1,2,3,4,5}
	//fmt.Printf("arr: %p\n",&arr)
	//sumArr(&arr)
	//fmt.Println(arr[0])
	findArr(5,arr)
}

func findArr(sum int, arr [5]int)  {
	for i := 0; i < len(arr); i++ {
		for j := i+1; j < len(arr); j++ {
			target := arr[i] + arr[j]
			if target == sum{
				fmt.Printf("(%d,%d)\n",i,j)
			}
		}
	}
}
