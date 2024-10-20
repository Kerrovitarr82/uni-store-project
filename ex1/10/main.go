package main

import (
	"fmt"
	"math"
	"strconv"
)

func isArmstrong(num int) bool {
	digits := strconv.Itoa(num)
	numDigits := len(digits)
	sum := 0

	for _, digit := range digits {
		d, _ := strconv.Atoi(string(digit))
		sum += int(math.Pow(float64(d), float64(numDigits)))
	}

	return sum == num
}

func main() {
	var num int

	fmt.Print("Введите число, которое надо проверить на принадлежность числам Армстронга: ")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}
	if isArmstrong(num) {
		fmt.Printf("%v - число Армстронга", num)
	} else {
		fmt.Printf("%v - НЕ число Армстронга", num)
	}
}
