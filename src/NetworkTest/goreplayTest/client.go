package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		req1 := &fasthttp.Request{}
		req1.Header.Add("X-Api-Key", "key-value")
		req1.SetRequestURI("http://127.0.0.1:8000/proxytest_header")
		resp1 := &fasthttp.Response{}
		fmt.Println(req1.Header.String())
		//client := &fasthttp.Client{}
		if err := fasthttp.Do(req1, resp1); err != nil {
			fmt.Println(err)
			fmt.Println("请求失败")
		}
		body := resp1.Body()
		fmt.Println(string(body))
		fmt.Println("time: " + time.Now().String())
	}
}
