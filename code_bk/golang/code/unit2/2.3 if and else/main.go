package main

import (
	"fmt"
)

func main() {
	if n := 3; n%2 == 0 {
		fmt.Println("Even Number")
	} else {
		fmt.Println("Odd Number")
	}
}
