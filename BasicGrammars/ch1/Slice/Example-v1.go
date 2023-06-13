package main

import (
	"fmt"
)

func main() {
	// 声明一个数组
	var fruits []string
	// 追加元素
	fruits = append(fruits, "apple", "banana", "orange")
	// 访问元素
	fmt.Println(fruits[0])
	// 修改元素
	fruits[1] = "grape"
	fmt.Println(fruits)

	//使用元素
	subSlice := fruits[1:3]
	fmt.Println(subSlice)

	// 使用make 函数创建切片并指定长度和容量

	number := make([]int, 3, 4)
	number[0] = 12345
	number[1] = 223456
	number[2] = 343343433
	fmt.Println(number)

	for index, value := range fruits {
		fmt.Println(index, value)
	}
}
