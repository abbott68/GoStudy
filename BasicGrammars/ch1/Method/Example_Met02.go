package main

import (
	"fmt"
	"math"
)

// 定义一个圆形类型
type Circle struct {
	radius float64
}

// 定义一个方法，计算圆形的面积
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// 定义一个方法，计算圆形的周长
func (c Circle) Circumference() float64 {
	return 2 * math.Pi * c.radius
}

// 定义一个方法，判断圆形是否为大圆
func (c Circle) IsBigCircle() bool {
	return c.radius > 10
}

func main() {
	// 创建一个圆形实例
	circle := Circle{radius: 5}

	// 调用圆形实例的方法
	area := circle.Area()
	fmt.Println("Area:", area) // 输出：Area: 78.53981633974483

	circumference := circle.Circumference()
	fmt.Println("Circumference:", circumference) // 输出：Circumference: 31.41592653589793

	isBigCircle := circle.IsBigCircle()
	fmt.Println("Is Big Circle:", isBigCircle) // 输出：Is Big Circle: false
}
