package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	person := newPerson("Feyoung", 21)
	fmt.Println(person)
	person = person.Dream(22)
	fmt.Println(person)
	admin := newPerson("admin", 12223)
	fmt.Println(admin)
}

func (p person) Dream(age int) *person {
	p.age = age
	return &p
}

func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age: age,
	}
}