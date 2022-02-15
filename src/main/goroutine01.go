package main

import (
	"fmt"
	"time"
)

/*
	无缓冲通道
		发送操作会阻塞，直到另一个 goroutine 在对应通道上执行接收操作。这是值传送完成，两个 goroutine 继续执行。
		同样，接收操作先执行，接收方 goroutine 将阻塞，直到另一个 goroutine 在同一通道上发送一个值。
*/
var x = make(chan int)

func sendGoroutine() {
	fmt.Println("准备发送, send goroutine sleep 10 second")
	time.Sleep(10 * time.Second)
	x <- 10

	fmt.Println("send goroutine continue")
}

func recvGoroutine() {
	fmt.Println("等待接收")
	value := <-x
	fmt.Printf("recv goroutine get value=%v", value)
}

func main() {
	go recvGoroutine()
	go sendGoroutine()

	time.Sleep(time.Minute * 10)
}
