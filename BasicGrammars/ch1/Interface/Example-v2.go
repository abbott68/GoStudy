package main

import (
	"fmt"
	"time"
)

// 定义接口
type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

func main() {
	var dilbert Employee

	dilbert.ID = 10
	dilbert.Name = "abbott"
	dilbert.DoB.UnixNano()
	dilbert.Position = "IT"
	dilbert.Salary = 11110
	dilbert.ManagerID = 5

	fmt.Printf("ID: %s\n ", dilbert.ManagerID)
	fmt.Printf("Name: %s\n ", dilbert.Name)
	fmt.Printf("DoB: %s\n ", dilbert.DoB)
	fmt.Printf("position: %s\n ", dilbert.Position)
	fmt.Printf("salary: %s\n", dilbert.Salary)
	fmt.Printf("ManagerID: %s\n", dilbert.ManagerID)

}
