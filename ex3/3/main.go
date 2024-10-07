package main

import (
	"fmt"
)

func reverseArr(arr [5]int) (rev [5]int) {
	n := len(arr)

	for i := 0; i < n; i++ {
		rev[i] = arr[n-1-i]
	}

	return rev
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
	fmt.Printf("%v - перевернутый массив\n", reverseArr(arr))
}
