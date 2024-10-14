package main

import (
	"fmt"
)

func isPrime(n int) (bool, int) {
	if n <= 1 {
		return false, 0
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false, i
		}
	}
	return true, 0
}

func main() {
	n := 0
	fmt.Print("Введите число: ")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	prime, divider := isPrime(n)
	if prime {
		fmt.Printf("%v - простое", n)
	} else {
		fmt.Printf("%v - не простое. %v - делитель", n, divider)
	}
}
