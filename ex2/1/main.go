package main

import (
	"errors"
	"fmt"
	"math"
)

func calculate(num1, num2 float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return num1 + num2, nil
	case "-":
		return num1 - num2, nil
	case "*":
		return num1 * num2, nil
	case "/":
		if num2 == 0 {
			return 0, errors.New("деление на ноль")
		}
		return num1 / num2, nil
	case "^":
		return num1 * num2, nil
	case "%":
		if num2 == 0 {
			return 0, errors.New("деление на ноль")
		}
		return math.Mod(num1, num2), nil
	default:
		return 0, errors.New("недопустимая операция")
	}
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Введите первое число: ")
	_, err := fmt.Scanf("%f", &num1)
	if err != nil {
		fmt.Println("Ошибка ввода первого числа:", err)
		return
	}

	fmt.Print("Введите второе число: ")
	_, err = fmt.Scanf("%f", &num2)
	if err != nil {
		fmt.Println("Ошибка ввода второго числа:", err)
		return
	}

	fmt.Print("Введите оператор (+, -, *, /, ^, %): ")
	_, err = fmt.Scanf("%s", &operator)
	if err != nil {
		fmt.Println("Ошибка ввода оператора:", err)
		return
	}

	result, err := calculate(num1, num2, operator)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Результат: %.2f\n", result)
}
