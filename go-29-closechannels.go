// 关闭 一个通道意味着不能再向这个通道发送值了。
// 这个特性可以用来给这个通道的接收方传达工作已经完成的信息。
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// 在接收的这个特殊的二值形式的值中， 如果 jobs 已经关闭了，并且通道中所有的值都已经接收 完毕，
			// 那么 more 的值将是 false。
			j, more := <- jobs
			if more {
				fmt.Println("received jobs", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<- done
}