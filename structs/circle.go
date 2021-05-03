package main

import "fmt"

type Circle struct {
	Point
	Raduis int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	w := Wheel{Circle{Point{6, 8}, 5}, 20}

	fmt.Printf("%#v\n", w)

	w.X = 42

	fmt.Printf("%#v\n", w)
}
