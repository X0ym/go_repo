package main

import (
	"bytes"
	"fmt"
	"github.com/valyala/fasthttp"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	request := fasthttp.AcquireRequest()
	response := fasthttp.AcquireResponse()

	request.SetRequestURI("http://127.0.0.1:8001/report")
	var buf bytes.Buffer
	request.SetBodyStream(&buf, -1)

	go func() {
		for i := 0; i < 10; i++ {
			code := getCode(10 * 1024)
			buf.Write([]byte(fmt.Sprintf("line:%d %s\r\n", i, code)))
		}
	}()

	err := fasthttp.Do(request, response)
	if err != nil {
		fmt.Println("请求失败：", err)
	}

	if response.Header.StatusCode() != http.StatusOK {
		fmt.Println("StatusCode=", response.Header.StatusCode())
	}
}

func getCode(codeLen int) string {
	// 1. 定义原始字符串
	rawStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890_"
	// 2. 定义一个buf，并且将buf交给bytes往buf中写数据
	buf := make([]byte, 0, codeLen)
	b := bytes.NewBuffer(buf)
	//随机从中获取
	rand.Seed(time.Now().UnixNano())
	for rawStrLen := len(rawStr); codeLen > 0; codeLen-- {
		randNum := rand.Intn(rawStrLen)
		b.WriteByte(rawStr[randNum])
	}
	return b.String()
}
