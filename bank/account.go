package main

import (
	"flag"
	"fmt"
)
import "sync"

var (
	mu      sync.Mutex // guards balance
	balance int
)

func Deposit(amount int) int {
	mu.Lock()
	balance = balance + amount
	fmt.Printf("Depositing: %v\n", balance)
	mu.Unlock()
	return balance
}

func Balance(amount int) int {
	mu.Lock()
	b := balance + amount
	mu.Unlock()
	return b
}

func main() {
	var dep = flag.Int("deposit", 0, "Enter amount to deposit")
	flag.Parse()
	b := Balance(1700)
	d := Deposit(*dep)

	fmt.Printf("Total in you account: %v\n", b+d)
}
