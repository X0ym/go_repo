package main

import "fmt"

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
	//var mu sync.Mutex
	//
	//mu.Lock()
	//
	//mu.Unlock()

	fmt.Println(a)
}
