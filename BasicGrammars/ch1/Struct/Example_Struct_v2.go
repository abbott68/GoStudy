package main

import "fmt"

// 定义关于Person地址结构体
type Address struct {
	City    string
	Country string
}

// 定义Person结构体
type Person01 struct {
	Name    string
	Age     int
	Address Address
}

//定义方法

func (p1 Person01) PrintInfo() {
	fmt.Printf("name:%s\n", p1.Name)
	fmt.Printf("age: %d\n", p1.Age)
	fmt.Printf("Address: %s, %s \n", p1.Address.Country, p1.Address.City)
}
func main() {
	//创建一个Person01机构体实例
	p := Person01{
		Name: "abbott",
		Age:  18,
		Address: Address{
			City:    "CN",
			Country: "Beijing",
		},
	}
	//调用Person01结构体的方法
	p.PrintInfo()
}
