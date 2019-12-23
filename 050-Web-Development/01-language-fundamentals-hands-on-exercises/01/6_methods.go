package main

import (
	"fmt"
	"math"
)

const pi = 3.14157

type square struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.breadth
}

func (c circle) area() float64 {
	return pi * math.Pow(c.radius, 2)
}

func main() {
	s1 := square{
		length:  10.0,
		breadth: 10.0,
	}

	fmt.Println("The area is: ", s1.area())

	c1 := circle{radius: 15.123}
	fmt.Println("The area is: ", c1.area())
}
