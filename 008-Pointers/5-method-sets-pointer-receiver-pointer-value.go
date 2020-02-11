/*

This method takes pointer receivers and pointer values in Golang

*/

package main

import (
	"fmt"
)

const pi = 3.14157

type circle struct {
	radius float32
}

func (c *circle) area() float32 {
	fmt.Println("The value of c is:", c)
	fmt.Println("The value of &c is:", &c)
	c.radius = 44
	fmt.Println("from the area function, after change, c is now:", c)
	return pi * c.radius * c.radius
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println("The area of circle is ", s.area())
}

func main() {
	c1 := circle{
		radius: 10,
	}
	fmt.Println("value of c1 from main before calling info is: ", c1)
	fmt.Println("address of c1 from main is before calling info: ", &c1)
	fmt.Println("Calling the info function..")
	info(&c1)
	fmt.Println("c1 now from main is: ", c1)
}
