package main

import "fmt"
import "math"
// 定义接口, 要在 Go 中实现一个接口，我们就需要实现接口中的所有 方法。
type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}
// 结构体类型 rect 实现了 geometry 接口
func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width+2*r.height
}
// 结构体类型 circle 实现了 geometry 接口
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2*math.Pi*c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width:3, height:4}
	c := circle{radius:5}
	measure(r)
	measure(c)
}