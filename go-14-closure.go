package main

import "fmt"

// func() int  表示匿名函数和返回值类型,
// intSeq函数返回一个匿名函数,而intSeq的变量i在匿名函数中返回
func intSeq() func() int{
	i := 0
	return func() int {
		i ++
		return i
	}
}

// 调用 intSeq 函数，将返回值（一个函数）赋给 nextInt变量。
// 这个函数的值包含了自己的值 i，这样在每 次调用 nextInt 时都会更新 i 的值。
func main() {
	// 此时变量i存储在nextInt这个闭包中
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}