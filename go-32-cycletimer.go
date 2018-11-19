package main

import "fmt"
import "time"

func main() {
	ticker := time.NewTicker(time.Millisecond*500)
	go func() {
		// 一个通道用来发送数据。 这里我们在这个通道上使用内置的 range 来迭代值每隔 500ms 发送一次的值。
		for t := range ticker.C {
			fmt.Println("Tick at", t)
		}
	}()

	time.Sleep(time.Millisecond*1600)
	// 打点器可以和定时器一样被停止。一旦一个打点停止了， 将不能再从它的通道中接收到值。
	ticker.Stop()
	fmt.Println("Ticker stopped")
}