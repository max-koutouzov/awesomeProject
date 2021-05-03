package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert *Employee

func EmployeeById(id int) *Employee {
	fmt.Println(EmployeeById(dilbert.ManagerID).Position)
	return EmployeeById(id)
}

func main() {
	EmployeeById(2)
}
