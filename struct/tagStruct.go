package main

import (
	"fmt"
	"reflect"
)

//Student 学生
type Students struct {
	ID     int      `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string   //json序列化是默认使用字段名作为key
	Name   []string //私有不能被json包访问
}

type XiaoMing struct {
	age int
	Students
}

func main() {
	var data *byte
	var in interface{}

	fmt.Println(data)
	fmt.Println(in)
	in = data

	fmt.Println(IsNil(in))
	fmt.Println(in == nil)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}