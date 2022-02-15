package main

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	data := "this is data"
	code, body, err := fasthttp.Post([]byte(data), "http://127.0.0.1:8001/", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Code: ", code)
	fmt.Println("body: ", string(body))
}
