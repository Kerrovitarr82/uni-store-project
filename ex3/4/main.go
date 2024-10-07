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
	n := 0
	fmt.Println("Введите число, до которого нужно найти простые числа.")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for i := 0; i <= n; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
