package main

import "fmt"

func main()  {
	s1 := test(func() string {
		return "100"
	})
	fmt.Println(s1)
}

func test(fn func() string) string {
	return fn()
}
