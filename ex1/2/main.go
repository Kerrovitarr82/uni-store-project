package main

import (
	"fmt"
)

func convertCelsius(temp int) int {
	return temp + 32
}

func convertFahrenheit(temp int) int {
	return temp - 32
}

func main() {
	choice := ""
	temp := 0
	fmt.Printf("Выберете вид конвертации.\n1) Из цельсия в фаренгейты\n2) Из фаренгейта в цельсии\n")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return
	}
	fmt.Println("Введите температуру")
	_, err = fmt.Scanln(&temp)
	if err != nil {
		return
	}
	switch choice {
	case "1":
		fmt.Printf("%v градусов Фаренгейта", convertCelsius(temp))
	case "2":
		fmt.Printf("%v градусов Цельсия", convertFahrenheit(temp))
	}
}
