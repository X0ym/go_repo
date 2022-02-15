package Test

import (
	"fmt"
	"time"
)

func GoroutineTest1() {
	ch1 := make(chan int)

	go receiveBefore(ch1)

	ch1 <- 1000

	time.Sleep(time.Millisecond * 100)
}

func receiveBefore (c chan int) {
	//time.Sleep(100 * time.Millisecond)
	x := <-c
	fmt.Println("channel recv before :", x)
}

func receiveBeforeNew (c chan int) {
	//time.Sleep(100 * time.Millisecond)
	x1 := <-c
	x2 := <- c
	x3 := <- c
	fmt.Println("channel recv 1 :", x1)
	fmt.Println("channel recv 2 :", x2)
	fmt.Println("channel recv 2 :", x3)
}

func receiveAfter(c chan int) {
	//time.Sleep(1000 * time.Millisecond)
	x := <- c
	fmt.Println("channel recv after :", x)
}

func GoroutineTest2() {
	ch1 := make(chan int, 3)

	go receiveBeforeNew(ch1)

	ch1 <- 10
	ch1 <- 20
	ch1 <- 30

	time.Sleep(time.Millisecond * 100)
}