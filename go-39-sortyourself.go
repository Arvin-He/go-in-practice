// 自定义排序
package main

import "sort"
import "fmt"

// 使用自定义函数进行排序，需要一个对应的 类型。
type ByLength []string

// 在ByLength类型中实现了 sort.Interface 的 Len，Less 和 Swap 方法，
// 这样我们就可以使用 sort 包的通用 Sort 方法了，
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
// Less 控制实际的自定义排序逻辑。
func (s ByLength) Less(i, j int) bool{
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}