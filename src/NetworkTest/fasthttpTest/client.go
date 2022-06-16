package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	req1 := &fasthttp.Request{}
	req1.SetRequestURI("http://127.0.0.1:8001/")
	resp1 := &fasthttp.Response{}

	//client := &fasthttp.Client{}
	if err := fasthttp.Do(req1, resp1); err != nil {
		fmt.Println(err)
		fmt.Println("请求失败")
	}
	body1 := resp1.Body()
	fmt.Println("body1: " + string(body1))

	// example2
	req2 := &fasthttp.Request{}
	req2.SetRequestURI("http://127.0.0.1:8001/")
	resp2 := &fasthttp.Response{}

	client := &fasthttp.Client{}
	if err := client.Do(req2, resp2); err != nil {
		fmt.Println(err)
		fmt.Println("请求失败")
	}
	body2 := resp2.Body()
	fmt.Println("body2: " + string(body2))

	// example3
	c := &fasthttp.Client{}
	code, body, err := c.Get(nil, "http://www.baidu.com")
	if err != nil {
		fmt.Println("请求百度失败")
	}
	if code != fasthttp.StatusOK {
		fmt.Println("响应失败")
	}
	fmt.Println(len(body))
	fmt.Println(string(body))
}
