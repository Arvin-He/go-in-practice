// Defer 被用来确保一个函数调用在程序执行结束前执行。同 样用来执行一些清理工作。 
// defer 用在像其他语言中的 ensure 和 finally用到的地方。

package main

import "fmt"
import "os"

func main() {
	f := createFile("f:/defer.txt")
	// 使用 defer 通过 closeFile 来关闭这个文件。
	//  这会在封闭函数 （main）结束时执行，就是 writeFile 结束后。
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File){
	fmt.Println("closing")
	f.Close()
}


