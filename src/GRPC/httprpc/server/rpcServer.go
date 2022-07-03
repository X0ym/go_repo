package main

import (
	"fmt"
	"go_awesomeProject/src/GRPC/httprpc/service"
	"net"
	"net/http"
	"net/rpc"
)

func main() {
	s := new(service.ServiceA)
	rpc.Register(s)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":9091")
	if err != nil {
		fmt.Println("监听失败。 err:", err)
	}
	http.Serve(listener, nil)
}
