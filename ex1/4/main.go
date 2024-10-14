package main

import "fmt"

func matrixMult() [10][10]int {
	var matrix [10][10]int
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			matrix[i][j] = 0
		}
	}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			matrix[i][j] = (i + 1) * (j + 1)
		}
	}
	return matrix
}

func main() {
	matrix := matrixMult()
	for _, row := range matrix {
		for _, col := range row {
			fmt.Print(col, " ")
		}
		fmt.Println()
	}
}
