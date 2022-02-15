package main

import (
	"fmt"
	"net"
	"time"
)

const (
	NUMS int = 2000
)

func main() {

	start := time.Now().UnixNano()
	//fmt.Println(start)

	//client to server
	// 请求建立连接
	//fmt.Println("client try connect server")
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("client connect to server failed. err:", err)
	}
	//fmt.Println("client connection success")

	for i := 0; i < NUMS; i++ {
		// send data to server
		reqData := "client 发送请求：tell me current time\n"
		_, err1 := conn.Write([]byte(reqData))
		if err1 != nil {
			fmt.Println("client failed to send reqData. err:", err1)
		}

		// 接收server的数据
		//fmt.Println(i, "waiting for server's data")
		buf := make([]byte, 128)
		_, err2 := conn.Read(buf)
		if err2 != nil {
			fmt.Println("client read failed. err:", err2)
		}
		//fmt.Println("client recv data from server :", string(buf[:n]))

	}

	end := time.Now().UnixNano()
	total := end - start
	fmt.Println("总时长：", total, "ns")
	fmt.Println("平均耗时：", total/int64(NUMS), "ns")

	conn.Close()
}
