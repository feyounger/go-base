package main

import "fmt"

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}
func main() {
	var slices []string
	slices = append(slices, "1")
}

