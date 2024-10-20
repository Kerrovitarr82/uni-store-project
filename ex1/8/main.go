package main

import (
	"fmt"
)

func palindromeNumber(num int) (bool, error) {
	if num < 0 {
		return false, fmt.Errorf("error: negative number")
	}
	original := num
	reversed := 0

	for num > 0 {
		lastDigit := num % 10
		reversed = reversed*10 + lastDigit
		num = num / 10
	}

	return original == reversed, nil
}

func main() {
	var n int
	fmt.Print("Введите число, которое надо проверить на палиндром: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	isPalindrome, err := palindromeNumber(n)
	if err != nil {
		fmt.Println(err)
	} else {
		if isPalindrome {
			fmt.Println("Число является палиндромом")
		} else {
			fmt.Println("Число НЕ является палиндромом")
		}
	}
}
