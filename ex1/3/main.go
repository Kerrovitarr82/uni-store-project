package main

import (
	"fmt"
	"math"
	"sort"
)

func sortByAbs(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return math.Abs(float64(arr[i])) < math.Abs(float64(arr[j]))
	})
}

func main() {
	var n int
	fmt.Print("Введите количество элементов массива: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Ошибка: необходимо ввести положительное целое число")
		return
	}

	arr := make([]int, n)

	fmt.Println("Введите элементы массива:")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			fmt.Println("Ошибка ввода. Убедитесь, что вы вводите целые числа.")
			return
		}
	}

	fmt.Println("Исходный массив:", arr)
	sortByAbs(arr)
	fmt.Println("Отсортированный массив по абсолютным значениям:", arr)
}
