package main

import "fmt"

// 定义一个矩形结构体
type Rectangle struct {
	width  float64
	height float64
}

// 定义一个方法，计算矩形的面积
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func main() {
	// 创建一个矩形实例
	rect := Rectangle{
		width:  4.2,
		height: 3.7,
	}

	// 调用矩形实例的方法来计算面积
	area := rect.Area()
	fmt.Println("Area:", area) // 输出：Area: 15.54
}
