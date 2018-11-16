package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	var value = n * fact(n-1) 
	fmt.Println(n, value)
	return value
}

func main() {
	fmt.Println(fact(7))
}