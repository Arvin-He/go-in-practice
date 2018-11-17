// 超时 对于一个连接外部资源，或者其它一些需要花费执行时间 的操作的程序而言是很重要的。
// 得益于通道和 select，在 Go 中实现超时操作是简洁而优雅的。

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second*2)
		c1 <- "result 1"
	}()
	// 使用 select 实现一个超时操作。res := <- c1 等待结果，<-Time.After 等待超时 时间 1 秒后发送的值。
	// 由于 select 默认处理第一个 已准备好的接收操作，如果这个操作超过了允许的 1 秒 的话，将会执行超时 case。
	select {
	case res := <- c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second*2)
		c2 <- "result 2"
	}()

	// 允许一个长一点的超时时间 3 秒，将会成功的从 c2 接收到值，并且打印出结果
	select {
	case res := <- c2:
		fmt.Println(res)
	case <-time.After(time.Second*3):
		fmt.Println("timeout 2")
	}
}

// 使用这个 select 超时方式，需要使用通道传递结果。
// 这对于 一般情况是个好的方式，因为其他重要的 Go 特性是基于通道和 select 的。
// 接下来我们就要看到两个例子：timer 和 ticker。