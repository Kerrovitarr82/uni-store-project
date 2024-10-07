package main

import (
	"fmt"
)

func whichAgeGroup(age int) string {
	switch {
	case age < 10:
		return "ребенок"
	case age >= 10 && age <= 19:
		return "подросток"
	case age > 19 && age < 70:
		return "взрослый"
	case age > 70:
		return "пожилой"
	default:
		return "неизвестно"
	}
}

func main() {
	x := 0
	fmt.Println("Ввведите возраст.")
	_, err := fmt.Scanf("%d", &x)
	if err != nil {
		return
	}
	fmt.Printf("Этот человек - %v.\n", whichAgeGroup(x))
}
