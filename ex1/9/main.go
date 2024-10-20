package main

import (
	"fmt"
)

func minMax(arr []int) (int, int, error) {
	if len(arr) == 0 {
		return 0, 0, fmt.Errorf("error: empty array")
	}
	if len(arr) == 1 {
		return arr[0], arr[0], nil
	}

	min := 0
	max := 0

	for i := 0; i < len(arr); i++ {
		if arr[i] < min {
			min = arr[i]
		}
		if arr[i] > max {
			max = arr[i]
		}
	}

	return min, max, nil
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

	fmt.Println("Введите элементы массива через пробел:")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			fmt.Println("Ошибка ввода. Убедитесь, что вы вводите целые числа.")
			return
		}
	}

	fmt.Println("-------------------------------\nИсходный массив: ", arr)
	min, max, err := minMax(arr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("-------------------------------\nМинимум = %v\nМаксимум = %v\n", min, max)
	}
}
