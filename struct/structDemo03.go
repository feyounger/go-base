package main

import "fmt"

type Teacher struct {
	name string
	age  int
}

type Person struct {
	name string
	age  int8
}

func (p *Person) SetAge2(newAge int8) {
	p.age = newAge
}

func NewPerson(name string,age int8) *Person  {
	return &Person{
		name: name,
		age: age,
	}
}

func main() {
	p1 := NewPerson("测试", 25)
	fmt.Println(p1.age)
	p1.SetAge2(30)
	fmt.Println(p1.age)
}


//func main() {
//	m := make(map[string]Teacher)
//	teachers := []Teacher{
//		{name: "Feyoung1", age: 10},
//		{name: "Feyoung2", age: 10},
//		{name: "Feyoung3", age: 10},
//	}
//	for _, teacher := range teachers {
//		m[teacher.name] = teacher
//	}
//
//	for key, value := range m {
//		fmt.Println("key:",key,"value: ",value)
//	}
//
//}