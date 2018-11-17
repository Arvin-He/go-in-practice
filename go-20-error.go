// Go 语言习惯的做法是使用一个独立、明确的返回值来传递错误信息。而不是使用异常
// 按照惯例，错误通常是最后一个返回值并且是 error 类 型，一个内建的接口。

package main

import "errors"
import "fmt"

func f1(arg int) (int, error) {
	if arg == 42 {
		// errors.New 构造一个使用给定的错误信息的基本 error 值。
		return -1, errors.New("Can't work with 42")
	}
	// 返回错误值为 nil 代表没有错误。
	return arg+3, nil
}

// 定义argError结构体
type argError struct {
	arg int
	prob string
}

// 定义argError的一个方法, 通过实现 Error 方法来自定义 error 类型
func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		// &argError 语法表示建立一个新的结构体，并提供了 arg 和 prob 这两个字段 的值。
		return -1, &argError{arg, "Can't work with it"}
	}
	return arg+3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		// 注意在 if 行内的错误检查代码，在 Go 中是一个普遍的用法。
		if r, e := f1(i); e!=nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}
	// 如果想在程序中使用一个自定义错误类型中的数据，需要通过类型断言来得到这个错误类型的实例
	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}

