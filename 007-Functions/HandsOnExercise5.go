/*

create a type SQUARE
create a type CIRCLE
attach a method to each that calculates AREA and returns it
circle area= π r 2
square area = L * W
create a type SHAPE that defines an interface as anything that has the AREA method
create a func INFO which takes type shape and then prints the area
create a value of type square
create a value of type circle
use func info to print the area of square
use func info to print the area of circle

*/

package main

import (
	"fmt"
)

const pi = 3.14157

type square struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

func (s square) area() {
	fmt.Println(s.length * s.breadth)
}

func (c circle) area() {
	fmt.Println(pi * (c.radius * c.radius))
}

type shape interface {
	area()
}

func info(s shape) {
	switch s.(type) {
	case square:
		fmt.Println("The shape is square")
		s.(square).area()
	case circle:
		fmt.Println("The shape is circle")
		s.(circle).area()
	}
}

func main() {

	square1 := square{length: 10.00,
		breadth: 10.00,
	}
	circle1 := circle{radius: 10.00}

	info(square1)
	info(circle1)
}
