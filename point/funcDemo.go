package main

import "fmt"

func main() {
	demo := new(Demo)
	fmt.Println(demo.Add())
	fmt.Println(demo.Add())
	fmt.Println(demo.Add())
	demo.Delete()
}

type Demo struct {
	i int
}

func (d *Demo) Add() int {
	d.i++
	return d.i
}

func (d *Demo) Delete() {
	fmt.Println(d.i)
}