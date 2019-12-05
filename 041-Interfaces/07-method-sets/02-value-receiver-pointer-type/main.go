package main

import (
	"fmt"
	"math"
)

const pi = 3.14157

type circle struct {
	radius float64
}

func (c circle) area() float64 { // it is a value receiver
	return pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Area: ", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}

	info(&c) // sending in pointer type
}
