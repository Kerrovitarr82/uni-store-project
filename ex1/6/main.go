package main

import (
	"fmt"
)

func reverseNumber(num int) int {
	reversed := 0

	for num != 0 {
		lastDigit := num % 10
		reversed = reversed*10 + lastDigit
		num /= 10
	}
	return reversed
}

func main() {
	var num int
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&num)
	if err != nil {
		return
	}

	fmt.Println("Перевернутое число:", reverseNumber(num))
}
