package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {
	// exec.Command 函数帮助我们创 建一个表示这个外部进程的对象。
	dateCmd := exec.Command("date")
	// .Output 是另一个帮助我们处理运行一个命令的常见情况 的函数，
	// 它等待命令运行完成，并收集命令的输出。
	// 如果没 有出错，dateOut 将获取到日期信息的字节
	dateOut, err := dateCmd.Output()

	if err != nil {
		panic(err)
	}
	// 打印 一些信息到标准输出流。
	fmt.Println(">date")
	fmt.Println(string(dateOut))
	grepCmd := exec.Command("grep", "hello")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()

	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

// 注意，在生成命令时，我们需要提供显式描述的命令和参数 数组，
// 而不能只传递一个命令行字符串。
// 如果你想使用一个 字符串生成一个完整的命令，
// 那么你可以使用 bash 命令 的 -c 选项：