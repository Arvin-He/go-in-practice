package main

import "fmt"
import "strings"
import "crypto/sha1"

func main() {
	s := "sha1 this string"

	h := sha1.New()
	// 写入要处理的字节。如果是一个字符串，需要使用 []byte(s) 来强制转换成字节数组。
	h.Write([]byte(s))
	// 这个用来得到最终的散列值的字符切片。
	// Sum 参数可以用来给现有的字符切片追加额外的字节切片：一般不需要。
	bs := h.Sum(nil)

	fmt.Println(s)
	fmt.Println(bs)
	// SHA1 值经常以 16 进制输出，例如在 git commit 中。
	// 使用 %x 来将散列结果格式化为 16 进制字符串。
	fmt.Printf("%x\n", bs)
}

// 计算 MD5 散列，引入 crypto/md5 并使用 md5.New() 方法。
// 注意，如果你需要密码学上的安全散列，你需要小心的研究一下 哈希强度。