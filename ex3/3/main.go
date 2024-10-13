package main

import (
	"fmt"
	"math"
	"strconv"
)

func printArmstrongNumbers(start, end int) {
	for num := start; num <= end; num++ {
		if isArmstrong(num) {
			fmt.Print(num, " ")
		}
	}
}

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
	var start, end int

	fmt.Print("Введите начальное значение диапазона: ")
	_, err := fmt.Scanf("%d", &start)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Print("Введите конечное значение диапазона: ")
	_, err = fmt.Scanf("%d", &end)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Println("Числа Армстронга в диапазоне от", start, "до", end, ":")
	printArmstrongNumbers(start, end)
}
