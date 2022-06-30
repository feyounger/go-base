package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func main() {
	sendMsg("azsdcasd")
}
const httpTimeoutSecond = time.Duration(30) * time.Second
type message string

type Response struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int64  `json:"errcode"`
}
func sendMsg(message)  (*Response, error){
	res := &Response{}
	pushURL := "https://oapi.dingtalk.com/robot/send?access_token=068bb8f58f178d2e517c4b6d078796f9cb0466412641f0a0cbd952bbb754b4bd"
	req, err := http.NewRequest(http.MethodPost,pushURL,strings.NewReader(`{"msgtype": "text","text": {"content": "starkwang"}}`))
	if err != nil {
		fmt.Println("01")
	}

	req.Header.Add("Accept-Charset","utf8")
	req.Header.Add("Content-Type","application/json")

	client := new(http.Client)

	client.Timeout = httpTimeoutSecond
	fmt.Println(req)
	resp ,err := client.Do(req)
	fmt.Println(resp)
	if err != nil {
		fmt.Println("02")
	}

	defer resp.Body.Close()

	resultByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("03")
	}
	err = json.Unmarshal(resultByte, &res)
	if err != nil {
		fmt.Errorf("unmarshal http response body from json error = %v", err)
	}
	if res.ErrCode != 0 {
		return res, fmt.Errorf("send message to dingtalk error = %s", res.ErrMsg)
	}

	fmt.Println("res----",res)

	return res, nil
}