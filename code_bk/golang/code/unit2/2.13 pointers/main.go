package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	i := 10
	j := &i
	i = 11
	*j = 20
	fmt.Println(*j, i)

	p := &Point{X: 10, Y: 20}
	(*p).X = 11
	p.X = 21
	fmt.Println(p)
}
