package main

import (
	"fmt"
	"image/color"
)

type Point struct {
	X, Y float64
}

func (p Point) ScaleBy(factor float64) {
}

func (p Point) Distance(q Point) float64 {
	return p.Distance(p)
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

func (p ColoredPoint) Distance(q Point) float64 {
	return p.Point.Distance(q)
}

func (p *ColoredPoint) ScaleBy(factor float64) {
	p.Point.ScaleBy(factor)
}

func main() {
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Distance(&Point{1, 1}), red}
	var q = ColoredPoint{&Point{5, 4}, blue}
	fmt.Println(p.Point)
	fmt.Println(q.Point)
}

func Distance(p *Point) *Point {
	return p
}
