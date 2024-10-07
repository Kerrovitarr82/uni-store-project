package main

import (
	"fmt"
)

func factorial(num int) int {
	final := 1
	for i := 1; i <= num; i++ {
		final *= i
	}
	return final
}

func main() {
	x := 0
	fmt.Println("Ввведите число.")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		return
	}
	fmt.Printf("%v - факториал %v\n", factorial(x), x)
}
