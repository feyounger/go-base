package hello

import (
	"fmt"
	"testing"
)

func TestStr(t *testing.T) {
	str := "Go语言"
	fmt.Println("str 长度:", len(str))
}
