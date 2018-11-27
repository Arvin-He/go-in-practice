package main

import "os"
import "fmt"

func main() {
	// os.Args 提供原始命令行参数访问功能, 第一个参数是该程序的路径
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]
	// 可以使用标准的索引位置方式取得单个参数的值。
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

// 实验命令行参数，最好先使用 go build 编译一个可执行 二进制文件