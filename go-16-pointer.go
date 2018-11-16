package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

// *int 参数，意味着它用了一 个 int指针
func zeroptr(iptr *int) {
	// *iptr 解引用这个指针， 从它内存地址得到这个地址对应的当前值。
	// 对一个解引用的指针赋值将会改变这个指针引用的真实地址的值。
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)
	
	zeroval(i)
	fmt.Println("zeroval:", i)

	// 通过 &i 语法来取得 i 的内存地址，即指向 i 的指针。
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)

}