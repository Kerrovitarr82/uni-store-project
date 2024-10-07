package main

import (
	"fmt"
)

func sum(num int) (final int) {
	final += num % 10
	num /= 10
	final += num % 10
	num /= 10
	final += num % 10
	num /= 10
	final += num % 10
	return final

}

func main() {
	num := 0
	fmt.Println("Введите 4-х значное число")
	_, err := fmt.Scanln(&num)
	if err != nil {
		return
	}
	fmt.Printf("Сумма %v\n", sum(num))
}
