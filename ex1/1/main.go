package main

import (
	"fmt"
	"math"
)

var keys = [...]string{
	"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J",
	"K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	"U", "V", "W", "X", "Y", "Z",
}

func Index(v string) int {
	for i, vs := range keys {
		if vs == v {
			return i
		}
	}
	return -1
}

func reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func switchNotation(num string, startNotation int, endNotation int) (final string) {
	buff := 0
	for i := 0; i < len(num); i++ {
		buff += Index(string(num[i])) * int(math.Pow(float64(startNotation), float64(len(num)-i-1)))
	}
	for buff >= endNotation {
		final += keys[buff%endNotation]
		buff /= endNotation
	}
	final += keys[buff]

	return reverse(final)

}

func main() {
	num := ""
	fmt.Println("Введите число")
	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}
	startNotation := 0
	fmt.Println("Введите систему счисления вашего числа")
	_, err = fmt.Scanln(&startNotation)
	if err != nil {
		return
	}
	endNotation := 0
	fmt.Println("Введите конечную систему счисления")
	_, err = fmt.Scanln(&endNotation)
	if err != nil {
		return
	}
	fmt.Printf("Итог %v\n", switchNotation(num, startNotation, endNotation))
}
