package main

import (
	"fmt"
	"math"
)

func distanceBetweenPoints(points [4]float64) float64 {
	return math.Sqrt(math.Pow(points[2]-points[0], 2) + math.Pow(points[3]-points[1], 2))

}

func main() {
	points := [4]float64{}
	fmt.Println("Введите значения точек в таком порядке: x1 y1 x2 y2")
	for i, _ := range points {
		_, err := fmt.Scanf("%f", &points[i])
		if err != nil {
			return
		}
	}

	fmt.Printf("Расстояние между точками = %v", distanceBetweenPoints(points))
}
