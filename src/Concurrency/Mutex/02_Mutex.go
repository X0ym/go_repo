package main

import (
	"sync"
)

/**

Unlock 方法可以被任意的 goroutine 调用释放锁，即使是没持有这个互斥锁的 goroutine，也可以进行这个操作。
这是因为，Mutex 本身并没有包含持有这把锁的 goroutine 的信息，所以，Unlock 也不会对此进行检查。

在使用 Mutex 的时候，必须要保证 goroutine 尽可能不去释放自己未持有的锁，一定要遵循“谁申请，谁释放”的原则。
在实践中，我们使用互斥锁的时候，一般都会在同一个方法中获取锁和释放锁。但是，如果临界区只是方法中的一部分，
为了尽快释放锁，还是应该第一时间调用 Unlock，而不是一直等到方法返回时才释放。


*/

func main() {
	var mu sync.Mutex
	a := 1
	mu.Lock()
	a++
	mu.Unlock()

	mu.Unlock()
}

type Foo struct {
	mu    sync.Mutex
	count int
}

// Bar 在方法中使用 Mutex 时，使用 defer 进行解锁
func (f *Foo) Bar() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.count < 1000 {
		f.count += 3
		return
	}
	f.count++
	return
}
