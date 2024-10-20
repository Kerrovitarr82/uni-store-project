package main

import (
	"fmt"
)

var memo = make(map[int]int)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	if val, found := memo[n]; found {
		return val
	}

	result := fibonacci(n-1) + fibonacci(n-2)

	memo[n] = result

	return result
}

func main() {
	var n int
	fmt.Print("Введите номер числа Фибоначчи: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}

	fmt.Printf("Число Фибоначчи для %d: %d\n", n, fibonacci(n))
}
