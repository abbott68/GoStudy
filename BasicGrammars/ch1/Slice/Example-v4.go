package main

import "fmt"

// 切片拼接：假设有两个整数切片，需要将它们合并为一个新的切片
func main() {
	numbers1 := []int{1, 2, 3}
	numbers2 := []int{4, 5, 6}

	result := append(numbers1, numbers2...)

	fmt.Println(result) // 输出：[1 2 3 4 5 6]
}
