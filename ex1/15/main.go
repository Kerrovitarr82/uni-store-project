package main

import (
	"fmt"
)

func intToRoman(num int) string {
	romanSymbols := []struct {
		value  int
		symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, rs := range romanSymbols {
		for num >= rs.value {
			roman += rs.symbol
			num -= rs.value
		}
	}

	return roman
}

func main() {
	var number int
	fmt.Print("Введите арабское число: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		return
	}

	if number <= 0 || number > 3999 {
		fmt.Println("Пожалуйста, введите число в диапазоне от 1 до 3999.")
		return
	}

	romanNumeral := intToRoman(number)
	fmt.Printf("Римское число для %d: %s\n", number, romanNumeral)
}
