package main

import (
	"fmt"
	"math"
)

func QuadraticRoots(a, b, c float64) (complex128, complex128) {
	if a == 0 {
		// В случае a == 0, это не квадратное уравнение
		panic("Коэффициент 'a' не может быть равен нулю")
	}

	D := b*b - 4*a*c

	if D > 0 {
		root1 := (-b + math.Sqrt(D)) / (2 * a)
		root2 := (-b - math.Sqrt(D)) / (2 * a)
		return complex(root1, 0), complex(root2, 0)
	}

	if D == 0 {
		root := -b / (2 * a)
		return complex(root, 0), complex(root, 0)
	}

	realPart := -b / (2 * a)
	imaginaryPart := math.Sqrt(-D) / (2 * a)
	return complex(realPart, imaginaryPart), complex(realPart, -imaginaryPart)
}

func main() {
	a := 0.0
	b := 0.0
	c := 0.0
	fmt.Println("Введите коэффициенты a, b, c (в таком же порядке через пробел)")
	_, err := fmt.Scanf("%f %f %f", &a, &b, &c)
	if err != nil {
		return
	}
	root1, root2 := QuadraticRoots(a, b, c)
	if imag(root1) != 0 {
		fmt.Printf("Комплексные корни уравнения: %v и %v\n", root1, root2)
	} else {
		fmt.Printf("Корни уравнения: %v и %v\n", real(root1), real(root2))
	}

}
