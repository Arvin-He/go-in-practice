package main

import "fmt"

func sum(args ... int) {
	fmt.Print(args, " ")
	total := 0
	for _, num := range args {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1,2, 3,4}
	sum(nums...)
}