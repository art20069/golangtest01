package main

import "fmt"

// bool

// string

// int  int8  int16  int32  int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte // alias for uint8

// rune // alias for int32

// float32 float64

// complex64 complex128


func main() {
	var flag bool
	var num int
	var sum int = 20
	// var age int = 18
	age := 18
	i, j, k := 1, 2, 3
	pi := 3.14
	fmt.Println(pi, uint(pi))
	fmt.Println(flag, num, sum, age, i, j, k)
}