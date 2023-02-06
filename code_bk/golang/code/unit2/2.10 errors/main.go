package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrInvalidRadius = errors.New("radius must be greater than 0")

func Area(r float64) (float64, error) {
	if r <= 0 {
		return 0.0, ErrInvalidRadius
	}
	return math.Pi * r * r, nil
}

func main() {
	if result, err := Area(0); errors.Is(err, ErrInvalidRadius) {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
