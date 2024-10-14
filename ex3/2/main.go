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
	start := 0
	end := 0
	fmt.Print("Введите начало диапазона простых чисел: ")
	_, err := fmt.Scanf("%d", &start)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	fmt.Print("Введите конец диапазона простых чисел: ")
	_, err = fmt.Scanf("%d", &end)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}
	for i := start; i <= end; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
}
