package main

import (
	"fmt"
)

func hasIntersection(a1, a2, b1, b2, c1, c2 float64) bool {
	if a1 > a2 {
		a1, a2 = a2, a1
	}
	if b1 > b2 {
		b1, b2 = b2, b1
	}
	if c1 > c2 {
		c1, c2 = c2, c1
	}

	maxStart := maxThree(a1, b1, c1)
	minEnd := minThree(a2, b2, c2)

	return maxStart <= minEnd
}

func maxThree(x, y, z float64) float64 {
	if x > y {
		if x > z {
			return x
		}
		return z
	}
	if y > z {
		return y
	}
	return z
}

func minThree(x, y, z float64) float64 {
	if x < y {
		if x < z {
			return x
		}
		return z
	}
	if y < z {
		return y
	}
	return z
}

func main() {
	var a1, a2, b1, b2, c1, c2 float64

	fmt.Print("Введите начальную и конечную точки первого отрезка: ")
	_, err := fmt.Scanf("%f %f", &a1, &a2)
	if err != nil {
		fmt.Println("Ошибка ввода первого отрезка:", err)
		return
	}

	fmt.Print("Введите начальную и конечную точки второго отрезка: ")
	_, err = fmt.Scanf("%f %f", &b1, &b2)
	if err != nil {
		fmt.Println("Ошибка ввода второго отрезка:", err)
		return
	}

	fmt.Print("Введите начальную и конечную точки третьего отрезка: ")
	_, err = fmt.Scanf("%f %f", &c1, &c2)
	if err != nil {
		fmt.Println("Ошибка ввода третьего отрезка:", err)
		return
	}

	if hasIntersection(a1, a2, b1, b2, c1, c2) {
		fmt.Println("Отрезки пересекаются")
	} else {
		fmt.Println("Отрезки не пересекаются")
	}
}
