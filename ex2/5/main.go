package main

import (
	"fmt"
)

func divisionOnFiveAndThree(num int) bool {
	if num%3 == 0 && num%5 == 0 {
		return true
	}
	return false
}

func main() {
	x := 0
	fmt.Println("Ввведите число.")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		return
	}
	if divisionOnFiveAndThree(x) {
		fmt.Printf("Делится.\n")
	} else {
		fmt.Printf("Не делится.\n")
	}
}
