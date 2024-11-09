package main

import (
	"fmt"
	"sort"
)

func sortArray(arr []int) []int {
	sort.Ints(arr)
	return arr
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

	sortedArr := sortArray(arr)
	fmt.Println("Отсортированный массив:", sortedArr)
}
