package main

import "fmt"

// 定义学生类型
type Student01 struct {
	name   string
	grades []int
}

// 定义方法，计算学生的平均成绩
func (s Student01) CalculateAverage() float64 {
	sum := 0
	for _, grade := range s.grades {
		sum += grade
	}
	average := float64(sum) / float64(len(s.grades))
	return average
}

func main() {
	//创建一个学生实例
	student01 := Student01{name: "abbott", grades: []int{80, 75, 90, 85, 95}}
	//调用学生示例的方法，计算平均成绩
	average := student01.CalculateAverage()
	fmt.Printf("Student: %s\nAverage Grade: %.2f\n", student01.name, average)
}
