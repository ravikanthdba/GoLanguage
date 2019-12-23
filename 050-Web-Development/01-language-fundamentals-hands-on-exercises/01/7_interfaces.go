package main

import (
	"fmt"
	"math"
)

const piValue = 3.14159

type squares struct {
	side float64
}

type rectangle struct {
	length  float64
	breadth float64
}

type triangle struct {
	base   float64
	height float64
}

type circles struct {
	radius float64
}

func (s squares) area() float64 {
	return math.Pow(s.side, 2)
}

func (r rectangle) area() float64 {
	return r.breadth * r.length
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (c circles) area() float64 {
	return piValue * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func calcArea(s shape) {
	switch s.(type) {
	case squares:
		fmt.Println("The shape is a square")
	case rectangle:
		fmt.Println("The shape is a rectangle")
	case triangle:
		fmt.Println("The shape is a triangle")
	case circles:
		fmt.Println("The shape is a circle")
	}

	fmt.Printf("The area is %f\n", s.area())
}

func main() {
	s1 := squares{
		side: 15,
	}
	calcArea(s1)

	r1 := rectangle{
		length:  12,
		breadth: 20,
	}
	calcArea(r1)

	t1 := triangle{
		base:   22,
		height: 25,
	}
	calcArea(t1)

	c1 := circles{radius: 15}
	calcArea(c1)
}
