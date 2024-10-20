package main

import (
	"fmt"
)

func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func digitalRoot(n int) int {
	if n < 10 {
		return n
	}
	return digitalRoot(sumOfDigits(n))
}

func main() {
	var number int
	fmt.Print("Введите число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}

	root := digitalRoot(number)
	fmt.Printf("Цифровой корень числа %d равен %d.\n", number, root)
}
