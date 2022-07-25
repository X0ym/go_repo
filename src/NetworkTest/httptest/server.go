package main

import (
	"fmt"
	"net/http"
)

/*

ServeMux is an HTTP request multiplexer.
It matches the URL of each incoming request against a list of registered
patterns and calls the handler for the pattern that
most closely matches the URL.

获取 ServeMux: NewServeMux 返回 ServeMux 指针

Handler 接口
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}




*/
func main() {
	//serverTest3()
	serverTest4()
}

func serverTest1() {
	/*
		使用默认 DefaultServeMux
	*/
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/test1", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("url:", request.URL.String())
	})
	err := http.ListenAndServe(":8001", serveMux)
	if err != nil {
		fmt.Println("监听8001失败。err=", err)
	}
}

func serverTest2() {
	serveMux := http.NewServeMux()
	serveMux.HandleFunc("/test2", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("url:", request.URL.String())
	})

	err := http.ListenAndServe(":8002", serveMux)
	if err != nil {
		fmt.Println("监听 8002 失败.err=", err)
	}
}

func serverTest3() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/test1", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("url:", req.URL.String())
	})
	server1 := &http.Server{Addr: ":8001", Handler: mux1}
	go server1.ListenAndServe()

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/test2", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("url:", req.URL.String())
	})
	server2 := &http.Server{Addr: ":8002", Handler: mux2}
	server2.ListenAndServe()
}

// ListenAndServe() 函数会阻塞住，开启多个 http 服务时，前面的 http 需要使用协程异步
func serverTest4() {
	mux1 := http.NewServeMux()
	mux1.HandleFunc("/test1", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("url:", req.URL.String())
	})
	//server1 := &http.Server{Addr: ":8001", Handler: mux1}
	//server.ListenAndServe()
	go http.ListenAndServe(":8001", mux1)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/test2", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("url:", req.URL.String())
	})
	//server2 := &http.Server{Addr: ":8002", Handler: mux2}
	//server2.ListenAndServe()
	http.ListenAndServe(":8002", mux2)
}
