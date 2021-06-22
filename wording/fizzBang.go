package main

import "fmt"

func main() {
	for i := 1; i <= 25; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("Odd number: %v\n", i)
			continue
		}
		if i%3 == 0 {
			fmt.Printf("Odd number: %v\n", i)
			continue
		}
		if i%5 == 0 {
			fmt.Printf("Divisable by 5: %v\n", i)
			continue
		}
	}
}
