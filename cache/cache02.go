package main

import (
	"fmt"
)

type User struct {
	Name string
}

func main() {
	fmt.Println("引用类型，直接赋值浅拷贝：")
	user1 := &User{
		Name: "why",
	}
	user2 := user1
	fmt.Println(*user1)
	fmt.Println(*user2)
	user1.Name = "jzm"
	fmt.Println(*user1)
	fmt.Println(*user2)
	fmt.Println("\n")

	fmt.Println("值类型，直接赋值深拷贝：")
	user3 := User{
		Name: "why",
	}
	user4 := user3
	fmt.Println(user3)
	fmt.Println(user4)
	user4.Name = "jzm"
	fmt.Println(user3)
	fmt.Println(user4)
	fmt.Println("\n")

	fmt.Println("引用类型，DeepCopy深拷贝：")
	user5 := &User{
		Name: "why",
	}
	user6 := new(User)
	fmt.Println(*user5)
	fmt.Println(*user6)
	DeepCopy(user6, user5)
	user6.Name = "jzm"
	fmt.Println(*user5)
	fmt.Println(*user6)
	fmt.Println("\n")

	fmt.Println("值类型，DeepCopy深拷贝：")
	user7 := User{
		Name: "why",
	}
	user8 := new(User)
	fmt.Println(user7)
	fmt.Println(*user8)
	DeepCopy(user8, user7)
	user8.Name = "jzm"
	fmt.Println(user7)
	fmt.Println(*user8)
	fmt.Println("\n")
}
