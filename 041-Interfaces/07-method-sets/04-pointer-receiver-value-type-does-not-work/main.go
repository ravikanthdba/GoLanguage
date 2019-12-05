package main

import (
	"fmt"
	"math"
)

const pi = 3.14157

type circle struct {
	radius float64
}

// receives a pointer
func (c *circle) area() float64 {
	return pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println("area:", s.area())
}

func main() {
	c := circle{
		radius: 5,
	}

	info(c) // value type being passed
}
