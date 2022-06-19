package main

import (
	"fmt"
	"sync"
)

/**

互斥锁的实现机制
	临界区：在并发编程中，被并发访问或修改的程序，就叫做临界区
	例如：临界区就是一个被共享的资源，或者说是一个整体的一组共享资源，
	比如对数据库的访问、对某一个共享数据结构的操作、对一个 I/O 设备的使用、对一个连接池中的连接的调用

	当很多线程同时访问临界区，就会出现资源竞争问题
	解决：互斥锁，限定临界区只能同时被一个线程持有



*/

const a = 1 << iota

func main() {
	counterTest2()
}

func counterTest1() {
	var count = 0
	// 使用 WaitGroup 等待 10 个 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 1000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

func counterTest2() {
	var count = 0
	var mu sync.Mutex
	// 使用 WaitGroup 等待 10 个 goroutine 完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 1000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

// 将 Mutex 嵌入到结构体中

type Counter1 struct {
	mu    sync.Mutex
	Count uint
}

type Counter2 struct {
	sync.Mutex
	Count uint
}

func counterTest3() {
	var counter Counter2
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count)
}
