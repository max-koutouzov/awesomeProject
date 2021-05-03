package main

import (
	"fmt"
	"math"
)

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func printMe(z string) string {
	return z
}

func main() {
	fmt.Println(printMe("Returning hypotenuse: "))
	fmt.Println(hypot(8, 11))
}
