package main

import (
	"fmt"
)

func sum(num [4]int) [4]int {
	num[0] *= 2
	num[1] *= 2
	num[2] *= 2
	num[3] *= 2
	return num

}

func main() {
	num := [4]int{0, 0, 0, 0}
	fmt.Println("Введите 4 числа через пробел")
	for i, _ := range num {
		_, err := fmt.Scanf("%d", &num[i])
		if err != nil {
			return
		}
	}

	fmt.Printf("Итоговый массив %v", sum(num))
}
