package main

import (
	"fmt"
	"math"
)

const pi = 3.14157

type Circle struct {
	radius float64
}

// We are passing the value receiver (there is no asteric symbol when we pass the receiver)
func (c Circle) area() float64 {
	return pi * math.Pow(c.radius, 2.0)
}

// Shape is an interface and any receiver implementing the method area() float64, is of type Shape
type Shape interface {
	area() float64
}

func info(s Shape) {
	fmt.Println("Area: ", s.area())
}

func main() {
	c1 := Circle{
		radius: 10,
	}

	fmt.Println(c1.area())

	c2 := Circle{
		radius: 5,
	}

	info(c2) // pass in a value type
}
