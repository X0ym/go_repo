package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func clockTest() {
	//
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal("err")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("获取连接失败")
		}

		handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	for {
		_, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
		if err != nil {
			fmt.Println("断开连接")
			return // 连接断开返回
		}
		time.Sleep(1 * time.Second)
	}
}
