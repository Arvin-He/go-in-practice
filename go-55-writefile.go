package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}


func main () {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("f:/go-write.txt", d1, 0644)
	check(err)

	f, err := os.Create("f:/dat2")
	check(err)
	// 打开文件后，习惯立即使用 defer 调用文件的 Close 操作。
	defer f.Close()
	// 写入字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3)
	// 调用 Sync 来将缓冲区的信息写入磁盘。
	f.Sync()
	// bufio 提供了和我们前面看到的带缓冲的读取器一样的带缓冲的写入器。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)
	// 使用 Flush 来确保所有缓存的操作已写入底层写入器。
	w.Flush()
}