package main

import (
	"fmt"
	"sync"
)

/**
Go并发
	Go语言的并发通过goroutine实现。goroutine类似于线程，属于用户态的线程
	goroutine是由Go语言的运行时（runtime）调度完成，而线程是由操作系统调度完成
	Go语言还提供channel在多个goroutine间进行通信
goroutine
	使用goroutine
		只需在调用函数的时候加上 go 关键字，就可以为函数创建一个 goroutine
		一个goroutine必定对应一个函数，可以创建多个goroutine去执行相同的函数
		如：go functionName
	goroutine调度
		GMP是Go语言运行时 (runtime) 层面的实现，在go语言层面实现的一套调度系统
			1. G 对应 goroutine
			2. P 管理一组 goroutine 队列，还会保存当前goroutine 运行的上下文信息（函数指针，堆栈地址及地址边界等信息）
			3. M（machine）是Go运行时（runtime）对操作系统内核线程的虚拟， M与内核线程一般是一一映射的关系， 一个goroutine最终是在M上执行
channel
	Go语言的并发模型是CSP（Communicating Sequential Processes），提倡通过通信共享内存而不是通过共享内存而实现通信。

	channel类型是一种引用类型
		声明 var varName chan 变量类型
		初始化  创建channel 使用 make 函数初始化，格式如：make(chan 变量类型 , [缓冲大小])
			eg 无缓冲 make(chan 元素类型)  有缓冲 make(chan 元素类型, size)

	channel支持的三种操作
		发送: ch <- x
		接收：y <- ch
			 <- ch
		关闭: close(ch)
*/

var wg sync.WaitGroup

func hello(i int) {
	fmt.Println("Hello goroutine", i)
	defer wg.Done()
}

func concurrentTest() {
	fmt.Println("Hello Goroutine!")

	// 声明 channel 变量
	ch1 := make(chan int, 10)
	fmt.Println(ch1)
}
