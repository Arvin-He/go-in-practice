package main
// 内置的 strconv 包提供了数字解析功能。
import "strconv"
import "fmt"

func main() {
	// 使用 ParseFloat 解析浮点数，这里的 64 表示解 析的数的位数。
	f, _ := strconv.ParseFloat("1.235", 64)
	fmt.Println(f)
	// 使用 ParseInt 解析整型数，参数 0 表示自动推断字符串所表示的数字的进制。
	// 64 表示返回的 整型数是以 64 位存储的。
	i, _ := strconv.ParseInt("-456", 0, 64)
	fmt.Println(i)
	// ParseInt 会自动识别出十六进制数。
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	u, _ := strconv.ParseUint("123", 0, 64)
	fmt.Println(u)
	// Atoi 是一个基础的 10 进制整型数转换函数。
	k, _ := strconv.Atoi("135")
	fmt.Println(k)
	// 在输入错误时，解析函数会返回一个错误。
	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}