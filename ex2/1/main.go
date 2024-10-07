package main

import (
	"fmt"
)

func evenOrNot(num int) bool {
	if num%2 == 0 {
		return true
	} else {
		return false
	}

}

func main() {
	num := 0
	fmt.Println("Ввведите число")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		return
	}
	if evenOrNot(num) {
		fmt.Println("Ваше число четное")
	} else {
		fmt.Println("Ваше число нечетное")
	}
}
