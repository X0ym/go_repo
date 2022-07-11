# Mutex

## 第一版实现

```go
package main

// CAS操作，当时还没有抽象出atomic包
    func cas(val *int32, old, new int32) bool
    func semacquire(*int32)
    func semrelease(*int32)
    // 互斥锁的结构，包含两个字段
    type Mutex struct {
        key  int32 // 锁是否被持有的标识
        sema int32 // 信号量专用，用以阻塞/唤醒goroutine
    }
  
    // 保证成功在val上增加delta的值
    func xadd(val *int32, delta int32) (new int32) {
        for {
            v := *val
            if cas(val, v, v+delta) {
                return v + delta
            }
        }
        panic("unreached")
    }
  
    // 请求锁
    func (m *Mutex) Lock() {
        if xadd(&m.key, 1) == 1 { //标识加1，如果等于1，成功获取到锁
            return
        }
        semacquire(&m.sema) // 否则阻塞等待
    }
  
    func (m *Mutex) Unlock() {
        if xadd(&m.key, -1) == 0 { // 将标识减去1，如果等于0，则没有其它等待者
            return
        }
        semrelease(&m.sema) // 唤醒其它阻塞的goroutine
    }   
```



## 第二版实现

## Mutex 的 4 种易错场景
1）Lock 和 Unlock 不是成对出现
场景1:没有 Unlock
场景2:Unlock 一个未加锁或已经解锁的 Mutex 而导致 panic(sync: unlock of unlocked mutex)

2)Copy 已使用的 Mutex
注意: sync 包的同步原语在使用后是不能复制的，因为 Mutex 是一个有状态的对象，其 state 字段记录了锁的状态
复制一个已经加锁的 Mutex 时，新的 Mutex 变量是已经加锁的状态（这显然不符合预期），且在此过程中 Mutex 的
也在变化

示例: 方法的传参为复制的方式传入的

```go
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	Count int
}

func main() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c) // 复制锁
}

// 这里Counter的参数是通过复制的方式传入的
func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}
```

3)重入
当一个 goroutine 成功获取到锁之后，如果其它线程再请求这个锁，就会处于阻塞等待的状态。但是，如果拥有这把锁
的线程再请求这把锁的话，就不会阻塞，而是成功返回，因此叫 可重入锁

可重入锁解决了代码重入和递归调用带来的死锁问题，并且可实现只有持有这个锁的 goroutine 才能 Unlock 这个锁

Mutex 是不可重入锁: 因为 Mutex 中没有记录哪个 goroutine 拥有这个锁

实现可重入锁
- 1 通过 hacker 的方式获取到 goroutine id，记录下获取锁的 goroutine id，它可以实现 Locker 接口。
- 2 调用 Lock/Unlock 方法时，由 goroutine 提供一个 token，用来标识它自己，而不是我们通过 hacker 
的方式获取到 goroutine id，但是，这样一来，就不满足 Locker 接口了。

获取 goroutine id
1) 可通过 runtime.Stack 方法获取堆栈信息，其中就包含了 goroutine id
```go
package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

func main() {
    GoId()
}

func GoId() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	// 得到id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
```
2)hacker 方式
获取运行时 goroutine 的指针，反解出对应的 goroutine 结构（每个运行的 goroutine 结构的 g 
指针保存在当前 goroutine 的 TLS 对象中

获取 goroutine id 的第三方库: petermattis/goid
```go
package main

import (
	"sync"
    "petermattis/goid"
)

func main() {
    
}

// RecursiveMutex 包装一个Mutex,实现可重入
type RecursiveMutex struct {
	sync.Mutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 这个goroutine 重入的次数
}
func (m *RecursiveMutex) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的goroutine就是这次调用的goroutine,说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}
	m.Mutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id,调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}
func (m *RecursiveMutex) Unlock() {
	gid := goid.Get()
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	if m.recursion != 0 { // 如果这个goroutine还没有完全释放，则直接返回
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.Mutex.Unlock()
}
```

