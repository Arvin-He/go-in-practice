package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name:"Alice", age:30})
	// 省略的字段将被初始化为零值。
	fmt.Println(person{name:"Freed"})
	fmt.Println(&person{name:"Ann", age: 40})

	s := person{name: "Sean", age:50}
	fmt.Println(s.name)

	// & 前缀生成一个结构体指针, 使用.来访问结构体字段。
	// 对结构体指针使用.指针会被自动解引用。
	sp := &s
	fmt.Println(sp.age)

	// 结构体是可变(mutable)的
	sp.age = 51
	fmt.Println(sp.age)

	// 初始化一个结构体元素时不指定字段名字,则按声明顺序初始化,否则报错
	// fmt.Println(person{20, "Bob"})
}

