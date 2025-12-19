package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance вычисляет расстояние от текущей точки до другой точки, используется формула: sqrt((x2-x1)^2 + (y2-y1)^2)
func (p *Point) Distance(other *Point) float64 {
	dx := other.x - p.x
	dy := other.y - p.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 0)
	point2 := NewPoint(3, 5)

	// Вычисляем расстояние между точками
	distance := point1.Distance(point2)

	fmt.Printf("Точка 1: (%.2f, %.2f)\n", point1.x, point1.y)
	fmt.Printf("Точка 2: (%.2f, %.2f)\n", point2.x, point2.y)
	fmt.Printf("Расстояние между точками: %.2f\n", distance)

}
