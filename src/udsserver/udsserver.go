package main

import (
	"fmt"
	"net"
	"os"
	"time"
)

const (
	NUMS int = 2000
)

const UnixSockPipePath = "/opt/yiming/unixsock_test.sock"

//const UnixSockPipePath = "/Users/xieym/mycode/my-code/goCode/awesomeProject/src/unixsock_test.sock"

func process(conn *net.UnixConn) {

	for i := 0; i < NUMS; i++ {
		buf := make([]byte, 512)
		// 读取到字节数组
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("uds : read from client failed. err:%", err)
		}
		//fmt.Println("uds : server 收到client的请求：", string(buf[:n]))

		// 响应
		conn.Write([]byte("uds_server响应数据: " + time.Now().String()))
		//fmt.Println("uds_server响应数据: " + time.Now().String())
		//fmt.Println(i)
	}
	fmt.Println("server over")
	conn.Close()
}

func main() {
	os.Remove(UnixSockPipePath)
	unixAddr, err := net.ResolveUnixAddr("unix", UnixSockPipePath)
	if err != nil {
		fmt.Println("invalid socket path")
	}

	// listen on the address
	unixListener, err := net.ListenUnix("unix", unixAddr)
	if err != nil {
		fmt.Println("uds : server listen failed. err:", err)
	}

	defer unixListener.Close()
	for {
		//fmt.Println("uds_server waiting for request")
		unixConn, err := unixListener.AcceptUnix()
		if err != nil {
			fmt.Println("uds_server failed to connect")
		}

		fmt.Println("server----等待处理完毕")
		go process(unixConn)
	}
}
