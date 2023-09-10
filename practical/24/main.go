package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func (p1 Point) DistanceTo(p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	point1 := NewPoint(1.5, 2.5)
	point2 := NewPoint(4.5, 6.5)
	distance := point1.DistanceTo(point2)
	fmt.Printf("Расстояние между точками: %.3f\n", distance)
}
