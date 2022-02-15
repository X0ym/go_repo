package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	r := strings.NewReader("this is body")
	response, err := http.Post("http://127.0.0.1:8001/", "text/plain", r)
	if err != nil {
		fmt.Println("post failed. err=", err.Error())
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("read resp failed. err=", err.Error())
	}
	fmt.Printf("resp body:%s", string(body))

}
