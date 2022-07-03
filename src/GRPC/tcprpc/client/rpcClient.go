package main

import (
	"fmt"
	"go_awesomeProject/src/GRPC/tcprpc/service"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:9091")
	if err != nil {
		fmt.Println("client dial failed. err:", err)
	}

	req := &service.GrpcRequest{X: 10, Y: 20}
	var reply service.GrpcReplay
	err = client.Call("ServiceA.Add", req, &reply)
	if err != nil {
		fmt.Println("client call failed. err:", err)
	}
	fmt.Println("data:", reply.Data)
	fmt.Println("msg:", reply.Msg)
	fmt.Println("code:", reply.Code)
}
