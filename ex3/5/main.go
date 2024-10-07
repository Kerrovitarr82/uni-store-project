package main

import (
	"fmt"
)

func arrSum(arr [5]int) (sum int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		sum += arr[i]
	}

	return sum
}

func main() {
	var arr [5]int
	fmt.Println("Введите числа массива размером 5 через пробел.")
	for i, _ := range arr {
		_, err := fmt.Scanf("%d", &arr[i])
		if err != nil {
			fmt.Printf(err.Error())
			return
		}
	}
	fmt.Printf("%v - сумма элементов массива\n", arrSum(arr))
}
