/*

Non pointer receiver works wiith pointer value

*/

package main

import (
	"fmt"
)

const pi = 3.14157;

type circle struct {
	radius float32;
}

func (c circle) area() float32 {
	fmt.Println("The value of c is ", c);
	fmt.Println("The address of c is ", &c);
	return pi * c.radius * c.radius;
}

type shape interface {
	area() float32;
}

func info(s shape) {
	fmt.Println("The value of s is ", s);
	fmt.Println("The address of s is ", &s);
	fmt.Println("The area is ", s.area());
}


func main() {

	c1 := circle {
		radius: 10,
	}

    fmt.Println("The value of c1 is ", c1);
    fmt.Println("The address of c1 is ", &c1);

	info(&c1);

}