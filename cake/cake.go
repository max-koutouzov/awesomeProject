package main

import "fmt"

type Cake struct {
	state  string
	amount int
}

func baker(cooked chan<- *Cake) {
	for {
		for i := 0; i < 100; i++ {
			fmt.Printf("Starting to cook cake #: %d", i)
			cake := new(Cake)
			cake.state = "cooked"
			cake.amount = i
			cooked <- cake // Baker is finished with this cake forever
		}
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // iser never touches this cake again
	}
}
