package main

import "fmt"

func Add(n, m int) int {
	return n + m
}

func Swap(n int, m int) (int, int) {
	return m, n
}

func main() {
	result1 := Add(10, 20)
	fmt.Println(result1)

	i, j := Swap(1, 2) // 2, 1
	fmt.Println(i, j)
}
