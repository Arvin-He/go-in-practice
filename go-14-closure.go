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

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}