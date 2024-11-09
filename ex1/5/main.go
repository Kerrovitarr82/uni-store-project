package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	if isPrime(n) {
		fmt.Println("Число простое.")
	} else {
		fmt.Println("Число не простое.")
	}
}
