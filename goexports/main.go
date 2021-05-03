package main

import (
	"fmt"
	"goexports/formatter"
	"goexports/math"
)

func main() {
	num := math.Double(2)
	output := formatter.Format(num)
	fmt.Println(output)
}