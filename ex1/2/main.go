package main

import (
	"fmt"
)

func euclideanAlgorithm(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var a, b int

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanf("%d", &b)
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	gcd := euclideanAlgorithm(a, b)

	fmt.Printf("Наибольший общий делитель %d и %d равен %d\n", a, b, gcd)
}
