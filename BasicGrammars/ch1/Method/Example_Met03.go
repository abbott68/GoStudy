package main

import "fmt"

// 定义学生类型
type Student struct {
	name  string
	grade int
}

// 定义一个方法，打印学生信息
func (s Student) PrintInfo() {
	fmt.Printf("Name: %s\nGrade: %d\n", s.name, s.grade)
}

// 定义一个方法，判断学生是否及格
func (s Student) IsPassing() bool {
	return s.grade >= 60
}

func main() {
	// 创建一个学生实例
	student := Student{name: "Alice", grade: 75}

	// 调用学生实例的方法
	student.PrintInfo() // 输出：Name: Alice, Grade: 75

	isPassing := student.IsPassing()
	fmt.Println("Is Passing:", isPassing) // 输出：Is Passing: true
}
