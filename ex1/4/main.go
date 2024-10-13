package main

import (
	"fmt"
	"sort"
)

func sortAndAppend(arr []int, arr2 []int) []int {
	sort.Ints(arr)
	sort.Ints(arr2)
	return append(arr, arr2...)
}

func main() {
	var n int
	fmt.Print("Введите количество элементов первого массива: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Ошибка: необходимо ввести положительное целое число")
		return
	}
	arr := make([]int, n)
	fmt.Println("Введите элементы первого массива:")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&arr[i])
		if err != nil {
			fmt.Println("Ошибка ввода. Убедитесь, что вы вводите целые числа.")
			return
		}
	}
	fmt.Print("Введите количество элементов второго массива: ")
	_, err = fmt.Scan(&n)
	if err != nil || n <= 0 {
		fmt.Println("Ошибка: необходимо ввести положительное целое число")
		return
	}
	arr2 := make([]int, n)
	fmt.Println("Введите элементы второго массива:")
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&arr2[i])
		if err != nil {
			fmt.Println("Ошибка ввода. Убедитесь, что вы вводите целые числа.")
			return
		}
	}

	fmt.Println("Исходные массивы:", arr, arr2)
	finalArr := sortAndAppend(arr, arr2)
	fmt.Println("Итоговый массив, собранный из двух отсортированных массивов:", finalArr)
}
