package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"
)

func main() {
	// 整数类型 Inter Types
	fmt.Println("=====打印整数类型=====")
	var age int = 30
	fmt.Println("年龄:", age)

	// 浮点数类型 Floating-Point Types
	fmt.Println("=====打印浮点数类型=====")
	var price float64 = 29.9
	fmt.Println("价格：", price)

	//字符串类型（String Type）
	fmt.Println("=====打印字符串类型=====")
	var name string = "abbott"
	fmt.Println("姓名：", name)

	// 布尔类型  Boolean Type
	fmt.Println("=====打印bool类型=====")
	var isStudent bool = true
	fmt.Println("是不是学生？", isStudent)

	// 复数类型 Complex Types
	fmt.Println("=====打印复数类型=====")
	var c complex128 = complex(3, 4)
	fmt.Println("复数:", c)

	// 字符类型 Rune Type
	fmt.Println("=====打印字符类型=====")
	var r rune = 'A'
	fmt.Println("字符", r)

	//错误类型 Error Type
	fmt.Println("=====打印错误类型=====")
	err := errors.NewBadRequest("发生了错误")
	fmt.Println("错误", err)

}
