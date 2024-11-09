package main

import (
	"fmt"
)

func triangleArea(base float64, height float64) float64 {
	return 0.5 * base * height
}

func main() {
	var base, height float64
	fmt.Print("Введите длину основания треугольника: ")
	fmt.Scan(&base)
	fmt.Print("Введите высоту треугольника: ")
	fmt.Scan(&height)

	area := triangleArea(base, height)
	fmt.Printf("Площадь треугольника: %.2f\n", area)
}
