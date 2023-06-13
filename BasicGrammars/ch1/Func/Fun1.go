package main

import "fmt"

// 定义一个函数，计算两个整数的和
func add(a, b int) int {
	return a + b
}

func main() {
	// 调用函数，并传入参数
	result := add(3, 5)

	// 打印函数返回的结果
	fmt.Println(result) // 输出：8
}
