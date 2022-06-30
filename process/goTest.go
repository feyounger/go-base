package main

import (
	"fmt"
	"io"
	"os"
)

type CustomReadWriter struct {
	name string
}

func (c CustomReadWriter) Read(p []byte) (n int, err error) {
	fmt.Println(p)
	return 1, nil
}

func (c CustomReadWriter) Write(p []byte) (n int, err error) {
	fmt.Println(p)
	return 1, nil
}

func main() {
	/*
		将f赋值给rw后，rw的动态类型是CustomReadWriter，接口类型是io.ReadWriter
		<io.ReadWriter, CustomReadWriter>
		<io.ReadWriter, *os.File>
		这两项都在itab缓存中，但是下面的断言明显不成功。
		所以，当断言成具体类型时，是不是应该直接检查rw的动态值的类型是不是与目标具体类型相等
	*/
	var rw io.ReadWriter
	f := CustomReadWriter{name: "poizon"}
	rw = f
	rf, ok := rw.(*os.File)
	fmt.Println(rf, ok)
}