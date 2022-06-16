package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://127.0.0.1:8001/test")
	if err != nil {
		fmt.Println("get failed. err=", err.Error())
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read resp failed. err=", err.Error())
	}
	fmt.Printf("resp body:%s", string(body))

}
