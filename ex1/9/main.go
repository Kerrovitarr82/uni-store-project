package main

import (
	"fmt"
)

func sumArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	var n int
	fmt.Print("Введите количество элементов в массиве: ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	result := sumArray(arr)
	fmt.Printf("Сумма элементов массива: %d\n", result)
}
