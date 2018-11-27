package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 对 os.Stdin 使用一个带缓冲的 scanner，直接使用方便的 Scan 方法来直接读取一行，
	// 每次调用 该方法可以让 scanner 读取下一行。
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Text 返回当前的 token
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}
	// 检查 Scan 的错误。文件结束符是可以接受的，并且 不会被 Scan 当作一个错误。
	if err := scanner.Err(); err!=nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}