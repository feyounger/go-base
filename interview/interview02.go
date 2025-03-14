package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

//sha256加密算法
func Sha256(message string, key string) (secretMessage []byte) {
	fmt.Println([]byte(key))
	//创建对应的sha256哈希加密算法,这里的key为加密密钥
	hash := hmac.New(sha256.New, []byte(key))
	// 写入加密数据，这里的"abc123"为加密信息，即需要加密的对象
	hash.Write([]byte(message))
	secretMessage = hash.Sum(nil)
	return
}

func main() {
	message := Sha256("20220903", "SDKbK1cChRisq9dcNFVc9NtwTUzO0iRlBxTAW41G3yg")
	fmt.Println(message)
}
