package main

import (
	"fmt"
	"net"
	"time"
)

const (
	NUMS int = 2000
)

//const UnixSockPipePath = "/Users/xieym/mycode/my-code/goCode/awesomeProject/src/unixsock_test.sock"
const UnixSockPipePath = "/opt/yiming/unixsock_test.sock"

func recvData(conn *net.UnixConn) string {
	buf := make([]byte, 512)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("client read failed. err:", err)
	}
	return string(buf[:n])
}

func main() {

	// 获取 UnixAddr 对象
	unixAddr, err := net.ResolveUnixAddr("unix", UnixSockPipePath)
	if err != nil {
		fmt.Println("invalid socket path")
	}

	start := time.Now().UnixNano()
	unixConn, err := net.DialUnix("unix", nil, unixAddr)
	if err != nil {
		fmt.Println("connect to server failed. err:", err)
	}

	for i := 0; i < NUMS; i++ {
		// send msg to server
		reqData := "client 发送请求： tell me current time\n"
		_, err1 := unixConn.Write([]byte(reqData))
		if err1 != nil {
			fmt.Println("client failed to send reqData. err:", err1)
		}
		//fmt.Println(reqData)
		// read data from server
		buf := make([]byte, 128)
		_, err2 := unixConn.Read(buf)
		if err != nil {
			fmt.Println("client read failed. err:", err2)
		}
		//fmt.Println("client recv data: ", string(buf[:n]))
		//fmt.Println(i)

	}
	end := time.Now().UnixNano()
	//fmt.Println(end)
	total := end - start
	fmt.Println("总耗时: ", total, "ns")
	fmt.Println("平均耗时: ", total/int64(NUMS), "ns")
	unixConn.Close()
}
