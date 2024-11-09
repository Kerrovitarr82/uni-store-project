package main

import (
	"fmt"
)

func findMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
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

	max := findMax(arr)
	fmt.Printf("Максимальный элемент в массиве: %d\n", max)
}
