package main

import "flag"
import "fmt"

// Go 提供了一个 flag 包，支持基本的命令行标志解析

func main() {
	wordPtr := flag.String("word", "foo", "a string")

	numbPtr := flag.Int("numb", 43, "an int")
	boolPtr := flag.Bool("fork", false, "a bool")
	// 用程序中已有的参数来声明一个标志也是可以的。
	// 注意: 在标志声明函数中需要使用该参数的指针。
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 所有标志都声明完成以后，调用 flag.Parse() 来执行 命令行解析。
	flag.Parse()


	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *boolPtr)
	fmt.Println("svar:", svar)
	fmt.Println("tail:", flag.Args())
}

// 尾随的位置参数可以出现在任何标志后面。 
// 注意，flag 包需要所有的标志出现位置参数之前
//  否则，这个标志将会被解析为位置参数
