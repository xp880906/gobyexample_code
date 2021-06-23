package main

import (
	"fmt"
	"math"
)

type Geometry interface {
	Area() float64
	Perim() float64
}

type Rect struct {
	Height, Width float64
}

func (r Rect) Area() float64 {
	return r.Height * r.Width
}

func (r Rect) Perim() float64 {
	return 2 * (r.Height + r.Width)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2.0)
}

func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}

func measure(g Geometry) {
	fmt.Println(g)
	fmt.Println("Area: ", g.Area())
	fmt.Println("Perim: ", g.Perim())
}

func main() {
	r := Rect{3, 4}
	c := Circle{5}

	measure(r)
	measure(c)
}
