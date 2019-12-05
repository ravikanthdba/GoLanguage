package main

import (
	"fmt"
	"math"
)

const pi = 3.14157

type circle struct {
	radius float64
}

// pointer receiver being received
func (c *circle) area() float64 {
	return pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("Value passed: ", s)
	fmt.Println("Area:", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}

	info(&c) // pointer value being passed
}
