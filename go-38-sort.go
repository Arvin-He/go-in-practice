package main

import "fmt"
import "sort"

func main() {
	// 排序方法是针对内置数据类型的；
	// 注意排序是原地更新的，所以他会改变给定的序列并且不返回 一个新序列。
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)
	// 检查一个序列是不是已经 是排好序的。
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}