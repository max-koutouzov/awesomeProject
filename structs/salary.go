package main

import "time"

func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}

func main() {
	AwardAnnualRaise(&Employee{
		ID:        0,
		Name:      "Dilbert",
		Address:   "555 Dilbo street",
		DoB:       time.Time{},
		Position:  "Techno guy",
		Salary:    10000,
		ManagerID: 0,
	})
}
