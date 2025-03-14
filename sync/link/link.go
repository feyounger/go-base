package link

import "fmt"
import _ "unsafe"

//go:linkname helloWorld go-base/sync/hello.Hello
func helloWorld() {
	fmt.Println("hello world")
}
