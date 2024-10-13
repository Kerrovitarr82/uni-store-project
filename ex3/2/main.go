package main

import (
	"fmt"
)

func fib(n int) []int {
	if n <= 0 {
		return []int{}
	}
	if n == 1 {
		return []int{0}
	}
	if n == 2 {
		return []int{0, 1}
	}

	final := make([]int, n)
	final[0] = 0
	final[1] = 1

	for i := 2; i < n; i++ {
		final[i] = final[i-1] + final[i-2]
	}

	return final
}

func main() {
	x := 0
	fmt.Println("Введите число.")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		return
	}
	fmt.Printf("%v - Числа Фибоначи до %v\n", fib(x), x)
}
