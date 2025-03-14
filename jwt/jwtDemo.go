package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

func Base64Decode(str string) string {
	reader := strings.NewReader(str)
	decoder := base64.NewDecoder(base64.RawStdEncoding, reader)
	// 以流式解码
	buf := make([]byte, 1024)
	// 保存解码后的数据
	dst := ""
	for {
		n, err := decoder.Read(buf)
		dst += string(buf[:n])
		if n == 0 || err != nil {
			break
		}
	}
	return dst
}

type Result struct {
	Jti string `json:"jti"`
	Exp int    `json:"exp"`
}

type GetCCEUserInfoResponse struct {
	OrgViewer         []string `json:"org_viewer"`         //字符串数组，拥有哪些组织的读权限
	Sub               string   `json:"sub"`                //用户ID
	OrgAdmin          []string `json:"org_admin"`          //字符串数组，拥有哪些组织的管理权限
	EmailVerified     bool     `json:"email_verified"`     //是否已经验证过邮箱
	Roles             []string `json:"roles"`              // 用户的系统角色
	Groups            []string `json:"groups"`             //表明用户在哪些组织与用户组下
	PreferredUsername string   `json:"preferred_username"` //系统内全局唯一的用户名
	Email             string   `json:"email"`              //邮箱
}

func main() {
	var res GetCCEUserInfoResponse
	str := "eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICItSHgtUDZvZ1pMeVFsQk1KS1dOVjF6U3QwWTNZMFozMjhVU2RfVlgxOGV3In0.eyJqdGkiOiJkYjEyMGQ3Ny1kMTU1LTRiMTAtOTk3NS0xNTM2ZDlhZDcyNmEiLCJleHAiOjE2NTkxNTEyNjYsIm5iZiI6MCwiaWF0IjoxNjU5MDY0ODY2LCJpc3MiOiJodHRwczovLzExOS4zLjE1OS4xOTUvYXV0aC9yZWFsbXMvQ0NFIiwiYXVkIjoiY2NlLWNsaWVudCIsInN1YiI6IjIxZmFiMmIxLTdlZmYtNGNhNC1hYzMxLWQyZjIwMWRhMzkwYiIsInR5cCI6IklEIiwiYXpwIjoiY2NlLWNsaWVudCIsImF1dGhfdGltZSI6MCwic2Vzc2lvbl9zdGF0ZSI6Ijk5NjZlMjczLWM1NTYtNGUyZS04MWEzLWVmZmY5ZTQwMmE0NCIsImFjciI6IjEiLCJvcmdfdmlld2VyIjpbIm9yZzEiXSwiZW1haWxfdmVyaWZpZWQiOmZhbHNlLCJyb2xlcyI6WyJiYXNpY191c2VyIl0sImdyb3VwcyI6WyIvb3JnMSIsIi9vcmcxL3VzZXItZ3JvdXAxIl0sInByZWZlcnJlZF91c2VybmFtZSI6InN1Yi11c2VyLTEiLCJlbWFpbCI6IjE3Nzc3Nzc3NzcxQG1haWwuY29tIn0.QKxItf1zSKgERBk6A326u6Ider68J7aBqaJPoRqpwcnpp-PsbMe_R9Zjq-E1-cCFERMnrbmSfONIY7QLFc2ubcyDhThaNvNrWKfJcFXXtBNcHo7ue3u6SAmrVyPMAhum2pbqzGVLYyj7zj6yNCFPtg_n1A5-fvbhxU6X-PZpA1H3RaU-2rP3DFX3JnM_mSIyynQcvHwJXDDrBie-di00qntDy6z7tMBbiRUAmoZ1R7DgKr8N8T3gW97pniaB10n-u5wRvjvzA7EVx4ZAHOcKJ7aE3euWio_1zdLKQea1Bw-lZ-BNX1Mt9tHQBw-SSJNugGc20YgbFvwICZMdaCoX7w"
	first := strings.Index(str, ".")
	second := strings.Index(str[first+1:], ".")
	strs := str[first+1 : first+second+1]
	//encodeStr := base64.StdEncoding.EncodeToString([]byte(str))
	//fmt.Println(encodeStr)
	decodeStr, err := base64.StdEncoding.DecodeString(strs)
	if err != nil {
		panic(err)
	}
	//fmt.Println(decodeStr)
	fmt.Println(string(decodeStr))
	json.Unmarshal(decodeStr, &res)
	fmt.Print(res)

}
