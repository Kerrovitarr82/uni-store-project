package main

import (
	"fmt"
)

func maxNum(x, y, z int) int {
	if x >= y && x >= z {
		return x
	} else if y >= x && y >= z {
		return y
	}
	return z
}

func main() {
	x := 0
	y := 0
	z := 0
	fmt.Println("Ввведите числа через пробел")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		return
	}
	_, err = fmt.Scanf("%d", &y)
	if err != nil {
		return
	}
	_, err = fmt.Scanf("%d", &z)
	if err != nil {
		return
	}
	fmt.Printf("%v - максимальное\n", maxNum(x, y, z))
}
