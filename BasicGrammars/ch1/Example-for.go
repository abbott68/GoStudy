package main

import "fmt"

/*
func main() {

		for i := 1; i <= 5; i++ {
			fmt.Println(i)
		}
		fmt.Println() // 输出空行

		// 示例2: 使用for循环计算1到5的和
		sum := 0
		for i := 1; i <= 5; i++ {
			sum += i
		}
		fmt.Println("Sum:", sum)

		fmt.Println() // 输出空行

		// 示例3: 使用for循环遍历字符串的字符
		message := "Hello, World!"
		for _, char := range message {
			fmt.Println(string(char))
		}
	}
*/
func main() {
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d", j, i, i*j)
		}
		fmt.Println()
	}
}
