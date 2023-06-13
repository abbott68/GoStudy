package main

import "fmt"

func main() {
	// 声明一个 3x3 的二维整数数组
	var matrix [3][3]int

	// 初始化二维数组元素
	matrix[0] = [3]int{1, 2, 3}
	matrix[1] = [3]int{4, 5, 6}
	matrix[2] = [3]int{7, 8, 9}

	// 访问二维数组元素
	fmt.Println(matrix[0][0]) // 输出：1
	fmt.Println(matrix[1][2]) // 输出：6

	// 修改二维数组元素
	matrix[2][1] = 88
	fmt.Println(matrix[2][1]) // 输出：88

	// 使用数组字面量初始化二维数组
	matrix2 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	fmt.Println(matrix2) // 输出：[[1 2] [3 4]]

	// 遍历二维数组元素
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Println(matrix[i][j])
		}
	}

	// 获取二维数组长度
	rows := len(matrix)
	cols := len(matrix[0])
	fmt.Println(rows, cols) // 输出：3 3
}
