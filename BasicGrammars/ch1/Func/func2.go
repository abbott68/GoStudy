package main

import "fmt"

// 定义一个函数，计算两个整数的和
func add1(a, b int) int {
	return a + b
}

// 定义一个函数，打印指定次数的欢迎消息
func greet(name string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println("Hello, " + name + "!")
	}
}

// 定义一个函数，计算两个浮点数的平均值
func average(nums ...float64) float64 {
	total := 0.0
	count := len(nums)

	for _, num := range nums {
		total += num
	}

	return total / float64(count)
}

func main() {
	// 调用函数并输出结果
	sum := add1(3, 5)
	fmt.Println("Sum:", sum) // 输出：Sum: 8

	greet("Alice", 5)
	/*
	   输出：
	   Hello, Alice!
	   Hello, Alice!
	   Hello, Alice!
	*/

	avg := average(2.5, 4.8, 6.1, 1.3, 9.7)
	fmt.Println("Average:", avg) // 输出：Average: 4.88
}
