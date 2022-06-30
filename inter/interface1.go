package main

import (
	"fmt"
	"reflect"
)

type People interface {
	Speak(string) string
}

type Student11 struct{
	Name string `json:"name"`
	Age int `json:"age"`

}

func (stu *Student11) Speak(think string) (talk string) {
	if think == "sb" {
		talk = "你是个大帅比"
	} else {
		talk = "您好"
	}
	return
}


func main() {
	h := Student11{
		Name: "小王",
		Age:  15,
	}
	retH := reflect.TypeOf(h)
	of := reflect.ValueOf(h)
	//获取结构体里的名称
	for i := 0; i < retH.NumField(); i++ {
		field := retH.Field(i)
		fmt.Println("结构体里的字段名",field.Name)
		fmt.Println("结构体里的字段属性:",field.Type)
		fmt.Println("结构体里面的字段的tag标签",field.Tag)
		fmt.Println(of.Field(i).Interface())
	}
	//提取tag标签里的信息
	name, bool := retH.FieldByName("Name")
	if bool {
		fmt.Println("tag标签的信息为",name.Tag.Get("json"))
	}
}