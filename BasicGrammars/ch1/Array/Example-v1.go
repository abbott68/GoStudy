package main

import (
	"fmt"
)

func main() {

	// 声明一个长度为5 的整数数组
	var nums [5]int
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50

	//访问数组中的元素
	fmt.Println(nums[0])
	fmt.Println(nums[2])
	fmt.Println(nums[1])
	fmt.Println(nums[3])
	fmt.Println(nums[4])

	//修改数组元素
	nums[0] = 21
	fmt.Println(nums[0])

	//使用数组字面量化数组
	nums1 := [3]int{1, 2, 3}
	fmt.Println(nums1)

	//使用for 循环变量数组
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	//获取数组长度
	length := len(nums)
	fmt.Println(length)
}
