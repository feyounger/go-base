package main

import "fmt"

func main() {
	//str := "Feyoungg"
	//println(len(str))
	//println(strings.Contains(str, "s"))
	//println(strings.HasPrefix(str, "F"))
	//println(strings.HasSuffix(str, "F"))
	//println(strings.Index(str, "g"))
	//println(strings.LastIndex(str, "g"))
	//arr := [...]struct{
	//	name string
	//	age int
	//}{
	//	{"user01",10},
	//	{"user02",9},
	//}
	//fmt.Println(arr)
	//
	//arr2 := [...][2]int{{1,1},{2,2},{3,3}}
	//var arr3  [3][2]int
	//fmt.Println(arr2)
	//fmt.Println(arr3)
	a := [2]int{}
	fmt.Printf("a: %p\n",&a)

	a = test(a)
	fmt.Println(a)
}

func test(x [2]int)(y [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
	fmt.Println(x)
	return x
}

