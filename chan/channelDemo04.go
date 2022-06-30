package main

import (
	"fmt"
	"time"
)

// DeferRecover defer recover from panic.
func DeferRecover(tag string, handlePanic func(error)) func() {
	return func() {
		if err := recover(); err != nil {
			fmt.Println(tag)
			if handlePanic != nil {
				handlePanic(fmt.Errorf("%v", err))
			}
		}
	}
}

// WithRecover recover from panic.
func WithRecover(tag string, f func(), handlePanic func(error)) {
	defer DeferRecover(tag, handlePanic)()
	f()
	panic("test")
}

// Go is a wrapper of goroutine with recover.
func Go(name string, f func(), handlePanic func(error)) {
	go WithRecover(fmt.Sprintf("goroutine %s", name), f, handlePanic)
}

func main() {
	Go("test-final", func() {
		fmt.Println("business")

	},func(err error){
		fmt.Println("recovery:",err)
	})

	time.Sleep(3*time.Millisecond)
}