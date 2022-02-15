package main

import (
	"fmt"
	"net"
	"time"
)

func process(conn *net.TCPConn) {
	for i := 0; i < 2000; i++ {
		buf := make([]byte, 512)
		// 读取到字节数组
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("tcp : read from client failed. err:%", err)
			return
		}
		//fmt.Println("tcp : server 收到client的数据：", string(buf[:n]))

		//fmt.Println("响应", i)
		//响应
		conn.Write([]byte("tcpserver响应数据: " + time.Now().String()))
	}

	conn.Close()
}

func main() {
	// 监听
	tcpAddr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:20000")
	//listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("tcp : server listen failed. err:", err)
		return
	}

	listen, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("tcp server listen failed. err: ", err)
	}

	for {
		//fmt.Println("tcp_server waiting for request")
		// 建立连接
		conn, err1 := listen.AcceptTCP()
		if err != nil {
			fmt.Println("tcp : server connection failed. err:", err1)
			return
		}
		conn.SetKeepAlive(true)
		//fmt.Println("server----连接成功")
		// 处理连接
		go process(conn)

		// 等待处理完连接，否则主线程结束导致连接断开
		//time.Sleep(1000000)
	}

}
