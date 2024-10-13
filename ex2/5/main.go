package main

import (
	"fmt"
)

func isLeapYear(year int) bool {
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				return true
			}
			return false
		}
		return true
	}
	return false
}

func main() {
	year := 0
	fmt.Println("Ввведите год")
	_, err := fmt.Scanf("%d", &year)
	if err != nil {
		return
	}
	if isLeapYear(year) {
		fmt.Printf("%v - високосный\n", year)
	} else {
		fmt.Printf("%v - невисокосный", year)
	}
}
