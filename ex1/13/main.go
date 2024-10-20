package main

import (
	"fmt"
	"time"
)

// Размеры поля
const (
	rows = 10
	cols = 10
)

func initializeField() [][]int {
	field := make([][]int, rows)
	for i := range field {
		field[i] = make([]int, cols)
	}

	field[1][2] = 1
	field[2][3] = 1
	field[3][1] = 1
	field[3][2] = 1
	field[3][3] = 1

	return field
}

func countAliveNeighbors(field [][]int, row int, col int) int {
	aliveNeighbors := 0
	directions := []struct{ x, y int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, dir := range directions {
		newRow, newCol := row+dir.x, col+dir.y
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			aliveNeighbors += field[newRow][newCol]
		}
	}

	return aliveNeighbors
}

func updateField(field [][]int) [][]int {
	newField := make([][]int, rows)
	for i := range newField {
		newField[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			aliveNeighbors := countAliveNeighbors(field, i, j)

			if field[i][j] == 1 { // Живая клетка
				if aliveNeighbors == 2 || aliveNeighbors == 3 {
					newField[i][j] = 1 // Остается живой
				}
			} else { // Мертвая клетка
				if aliveNeighbors == 3 {
					newField[i][j] = 1 // Оживает
				}
			}
		}
	}

	return newField
}

func printField(field [][]int) {
	for _, row := range field {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("█ ") // Живая клетка
			} else {
				fmt.Print(". ") // Мертвая клетка
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
	field := initializeField()

	for {
		printField(field)
		field = updateField(field)
		time.Sleep(500 * time.Millisecond) // Задержка между итерациями
	}
}
