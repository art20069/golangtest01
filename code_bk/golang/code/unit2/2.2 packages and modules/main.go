package main

import (
	"fmt"
	"course-go/shape"
)

func main() {
	area := shape.GetCircleArea(10)
	fmt.Println("Circle Area =", area)
}