package main

import "fmt"

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// 一个非空的通道也是可以关闭的， 但是通道中剩下的值仍然可以被接收到。
	close(queue)
	//  range 迭代从 queue 中得到的每个值。因为我们 在前面 close 了这个通道，
	// 这个迭代会在接收完 2 个值 之后结束。如果我们没有 close 它，
	// 我们将在这个循环中 继续阻塞执行，等待接收第三个值
	for elem := range queue {
		fmt.Println(elem)
	}
}

