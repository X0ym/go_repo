package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Go")
}

func main() {
	//   向DefaultServerMux添加处理器
	http.HandleFunc("/hello", sayHello)
	// 指定端口，和处理器，nil表示使用包变量 DefaultServerMux 作为处理器
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
	}
}
