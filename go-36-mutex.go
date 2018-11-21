package main

import "fmt"
import "math/rand"
import "runtime"
import "sync"
import "sync/atomic"
import "time"

func main() {
	var state = make(map[int]int)
	// mutex 将同步对 state 的访问
	var mutex = &sync.Mutex{}
	// ops记录对 state 的操作次数
	var ops int64 = 0
	// 运行 100 个 Go 协程来重复读取 state。
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			// 每次循环读取，使用一个键来进行访问， 
			// Lock() 这个 mutex 来确保对 state 的 独占访问，读取选定的键的值，Unlock() 这个 mutex，并且 ops 值加 1。
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				// 为了确保这个 Go 协程不会在调度中饿死，我们 在每次操作后明确的使用 runtime.Gosched() 进行释放。
				runtime.Gosched()
			}
		}()
	}
	// 运行 10 个 Go 协程来模拟写入操作，使用 和读取相同的模式。
	for w:=0; w<10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)	
				runtime.Gosched()
			}
		}()
	}
	// 让这 10 个 Go 协程对 state 和 mutex 的操作 运行 1 s。
	time.Sleep(time.Second)
	// 获取并输出最终的操作计数。
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	// 对 state 使用一个最终的锁，显示它是如何结束的
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}