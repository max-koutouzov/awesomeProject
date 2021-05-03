package main

import "fmt"

type Currency int

func currency() {
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)

	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}

	for i, v := range symbol {
		fmt.Printf("%d %v\n", i, v)
	}

}

func main() {
	currency()

}
