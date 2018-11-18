// 常规的通过通道发送和接收数据是阻塞的。
// 然而，可以使用带一个 default 子句的 select 来实现非阻塞 的发送、接收，
// 甚至是非阻塞的多路 select。
package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <- messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	//在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器。
	select {
	case msg := <- messages:
		fmt.Println("received message", msg)
	case sig := <- signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}