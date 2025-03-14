package main

import (
	"fmt"
	"unsafe"
)

type Person interface {
	growUp()
	growUp1()
}

type Student struct {
	age int
}

func (p Student) growUp() {
	p.age += 1
	return
}

func (p Student) growUp1() {
	p.age += 1
	return
}

func main() {
	var qcrao = Person(Student{age: 234243})

	fmt.Println(qcrao)
	iface := (*iface)(unsafe.Pointer(&qcrao))
	fmt.Printf("iface.tab.hash = %#x\n", iface.tab.hash)
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}
type itab struct {
	inter uintptr
	_type uintptr
	link  uintptr
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}
