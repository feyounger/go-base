package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	r, _ := os.Open("D:\\GoProject\\go-base\\text.txt")
	w = r
	rw, ok := w.(io.ReadWriter)
	fmt.Println(rw, ok)

}
