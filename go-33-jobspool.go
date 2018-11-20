package main

import "fmt"
import "time"

func worker(id int, jobs <- chan int, results chan <- int) {
	for j:= range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j*2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	// 启动了 3 个 worker，初始是阻塞的，因为 还没有传递任务。
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// 发送 9 个 jobs，然后 close 这些通道 来表示这些就是所有的任务了。
	for j :=1; j <=9; j++ {
		jobs <- j;
	}
	close(jobs)
	// 收集所有这些任务的返回值。
	for a := 1; a <=9; a++ {
		<- results
	}
}

//  9 个任务被多个 worker 执行。整个程序 处理所有的任务仅执行了 3s多 而不是 9s，是因为 3 个 worker 是并行的。
