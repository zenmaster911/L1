package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func newPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

func (p Point) distance(point Point) float64 {
	return math.Sqrt(math.Pow(p.x-point.x, 2) + math.Pow(p.y-point.y, 2))
}

func main() {
	a := newPoint(2, 2)
	b := newPoint(5, 2)
	fmt.Println(a.distance(*b))
}
