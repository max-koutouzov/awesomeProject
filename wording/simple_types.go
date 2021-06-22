package main

import (
	"fmt"
)

type Adder struct {
	start int
}

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee
	Reports []Employee
}

func (m Manager) FindNewEmployees() []Employee {
	fmt.Sprintf("%s", m.Reports)
	return m.Reports
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	myAddr := Adder{start: 10}
	f1 := myAddr.AddTo
	fmt.Println(myAddr.AddTo(5))
	fmt.Println(f1(10))

	s := Manager{
		Employee: Employee{
			Name: "Bob Bobson",
			ID:   "12345",
		},
		Reports: []Employee{},
	}
	fmt.Println(s.ID)
	fmt.Println(s.Description())
	fmt.Println(s.Reports)

}
