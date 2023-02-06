package main

import "fmt"

func main() {
	var num1 [2]int
	num1[0] = 1
	num1[1] = 2
	fmt.Println(len(num1), num1[0], num1[1])

	num2 := [3]int{1, 2, 3}
	fmt.Println(len(num2), num2)

	var num3 []int // nil
	if num3 == nil {
		fmt.Println("nil slice")
	}

	num3 = append(num3, 1)
	num3 = append(num3, 2)
	num3 = append(num3, 3)
	num3 = append(num3, 4)
	fmt.Println(num3)

	num4 := []int{1, 2, 3, 4}
	fmt.Println(len(num4), num4)

	num5 := num4[1:3]
	fmt.Println(num5)
	num6 := num4[1:]
	fmt.Println(num6)
	num7 := num4[:3]
	fmt.Println(num7)

	num8 := []int{1, 2, 3, 4}
	num9 := num8[1:3] // [2, 3]
	fmt.Println(num8) // [1, 2, 3, 4]
	num9[0] = 20
	fmt.Println(num8) // [1, 20, 3, 4]
	num8[2] = 30
	fmt.Println(num8) // [1, 20, 30, 4]
	fmt.Println(num9) // [20, 30]

	for i := range num8 {
		fmt.Println(i)
	}

	for i, v := range num8 {
		fmt.Println(i, v)
	}

	for _, v := range num8 {
		fmt.Println(v)
	}
}
