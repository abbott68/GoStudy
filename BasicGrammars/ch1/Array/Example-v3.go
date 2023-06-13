package main

import "fmt"

func main() {
	// 声明一个 2x3x4 的三维整数数组
	var cube [2][3][4]int

	// 初始化三维数组元素
	cube[0] = [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	cube[1] = [3][4]int{
		{13, 14, 15, 16},
		{17, 18, 19, 20},
		{21, 22, 23, 24},
	}

	// 访问三维数组元素
	fmt.Println(cube[0][0][0]) // 输出：1
	fmt.Println(cube[1][2][3]) // 输出：24

	// 修改三维数组元素
	cube[0][1][2] = 99
	fmt.Println(cube[0][1][2]) // 输出：99

	// 使用数组字面量初始化三维数组
	cube2 := [2][2][2]int{
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}
	fmt.Println(cube2) // 输出：[[[1 2] [3 4]] [[5 6] [7 8]]]

	// 遍历三维数组元素
	for i := 0; i < len(cube); i++ {
		for j := 0; j < len(cube[i]); j++ {
			for k := 0; k < len(cube[i][j]); k++ {
				fmt.Println(cube[i][j][k])
			}
		}
	}

	// 获取三维数组长度
	layers := len(cube)
	rows := len(cube[0])
	cols := len(cube[0][0])
	fmt.Println(layers, rows, cols) // 输出：2 3 4
}
