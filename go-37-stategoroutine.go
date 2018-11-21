// 用互斥锁进行了明确的锁定来让共享的 state 跨多个 Go 协程同步访问。
// 另一个选择是使用内置的 Go 协程和通道的同步特性来达到同样的效果。
package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main() {
	var readOps uint64 = 0
	var writeOps uint64 = 0
	// reads 和 writes 通道分别将被其他 Go 协程用来发 布读和写请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// state是被这个状态协程私有的。这个 Go 协程 反复响应到达的请求。
	// 先响应到达的请求，然后返回一个值到 响应通道 resp 来表示操作成功（或者是 reads 中请求的值）
	go func() {
		// state 被一个单独的 Go 协程拥有, 能够保证数据在并行读取时不会混乱。
		var state = make(map[int]int)
		for {
			select {
			case read := <- reads:
				read.resp <- state[read.key]
			case write := <- writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()
	// 启动 100 个 Go 协程通过 reads 通道发起对 state 所有者 Go 协程的读取请求。
	// 每个读取请求需要构造一个 readOp， 发送它到 reads 通道中，并通过给定的 resp 通道接收 结果。
	for r := 0; r < 100; r ++ {
		go func() {
			for {
				read := &readOp{
					key: rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<- read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 用相同的方法启动 10 个写操作。
	for w := 0; w<10; w++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	// 让 Go 协程们跑 1s。
	time.Sleep(time.Second)
	// 获取并报告 ops 值。
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}