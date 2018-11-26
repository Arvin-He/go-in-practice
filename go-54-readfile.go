package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func check(e error){
	if e != nil {
		panic(e)
	}
}

func main () {
	// 将文件内容读取到 内存
	dat, err := ioutil.ReadFile("f:/defer.txt")
	check(err)
	fmt.Println(string(dat))
	// 使用 os.Open 打开一个文件获取一个 os.File 值开始。
	f, err := os.Open("f:/defer.txt")
	check(err)
	// 最多读取 5 个字 节
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1))
	// Seek 到一个文件中已知的位置并从这个位置开 始进行读取。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	// Seek 到一个文件中已知的位置并从这个位置开 始进行读取。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	// 使用 ReadAtLeast 得到一个更 健壮的实现。
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	// 没有内置的回转支持，但是使用 Seek(0, 0) 实现。
	_, err = f.Seek(0, 0)
	check(err)
	// bufio 包实现了带缓冲的读取，这不仅对于很多小的读 取操作能够提升性能，也提供了很多附加的读取函数。
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))
	// 任务结束后要关闭这个文件（通常这个操作应该在 Open 操作后立即使用 defer 来完成）。
	f.Close()
}