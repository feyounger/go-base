package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/patrickmn/go-cache"
	"time"
)

func DeepCopy(dst, src interface{}) error {

	var buf bytes.Buffer

	if err := gob.NewEncoder(&buf).Encode(src); err != nil {

		return err

	}

	return gob.NewDecoder(bytes.NewBuffer(buf.Bytes())).Decode(dst)

}

// DefaultToken 设置超时时间和清理时间
var DefaultToken = cache.New(30*time.Minute, 30*time.Minute)

type Domain struct {
	Id   string
	Name string
}

func main() {
	var domains []Domain
	var result []Domain
	domains = append(domains, Domain{Id: "111", Name: "111"})
	domains = append(domains, Domain{Id: "222", Name: "222"})
	domains = append(domains, Domain{Id: "333", Name: "333"})
	DefaultToken.Add("domain", domains, time.Second*10)

	cache, ok := DefaultToken.Get("domain")
	if !ok {
		panic("无")
	}
	value := cache.([]Domain)
	//data, err := json.Marshal(value)
	//if err != nil {
	//	panic(err)
	//}
	//json.Unmarshal(data, &result)
	DeepCopy(result, value)
	fmt.Println("第一次取缓存：", result)

	result[0].Id = "1"
	result[0].Name = "1"

	cache, ok = DefaultToken.Get("domain")
	if !ok {
		panic("无")
	}
	result = cache.([]Domain)
	fmt.Println("第二次取缓存：", result)
}
