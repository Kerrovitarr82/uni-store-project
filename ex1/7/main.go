package main

import (
	"fmt"
)

func pascalTriangle(n int) ([][]int, error) {
	triangle := make([][]int, n)
	if n <= 0 {
		return nil, fmt.Errorf("глубина должна быть больше нуля")
	}

	for i := 0; i < n; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0] = 1
		triangle[i][i] = 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

	}
	return triangle, nil
}

func main() {
	var n int
	fmt.Print("Введите глубину треугольника Паскаля: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		return
	}
	triangle, err := pascalTriangle(n)
	if err != nil {
		fmt.Println(err)
	} else {
		for _, row := range triangle {
			for _, col := range row {
				fmt.Print(col, " ")
			}
			fmt.Println()
		}
	}
}
