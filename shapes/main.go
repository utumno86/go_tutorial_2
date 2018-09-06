package main

import (
	"fmt"
	"math"
)

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

type shape interface {
	getArea() float64
}

func main() {

	sq := square{
		sideLength: 2.5,
	}

	tri := triangle{
		height: 3,
		base:   4,
	}

	printArea(sq)
	printArea(tri)
}

func printArea(s shape) {
	fmt.Println("The area of the shape is", s.getArea())
}

func (s square) getArea() float64 {
	return math.Pow(s.sideLength, 2)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}
