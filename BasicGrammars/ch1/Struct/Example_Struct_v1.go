package main

import "fmt"

//定义人类型

type Person struct {
	Name string
	Age  int
}

func main() {
	// 创建一个person结构体实例
	p := Person{
		Name: "abbott",
		Age:  25,
	}
	//访问结构体的字断
	fmt.Println("name:", p.Name)
	fmt.Println("age:", p.Age)

	//修改结构体的字断值
	p.Age = 30
	fmt.Println("update age:", p.Age)
}
