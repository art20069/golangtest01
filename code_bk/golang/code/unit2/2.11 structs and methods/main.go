package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Point struct {
	X int
	Y int
}

func main() {
	c := Circle{10}
	a := c.Area()
	fmt.Println(c, a)

	p := Point{Y: 10, X: 20}
	fmt.Println(p)
}
