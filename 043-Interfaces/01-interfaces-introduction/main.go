package main

import "fmt"

type Square struct {
	size float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	height  float64
	breadth float64
}

func (s Square) area() float64 {
	return s.size * s.size
}

func (c Circle) area() float64 {
	return 3.14147 * c.radius * c.radius
}

func (t Triangle) area() float64 {
	return (0.5 * t.height * t.breadth)
}

type Shape interface {
	area() float64
}

func display(r Shape) {
	fmt.Printf("The type of the shape is %T\n", r)
	switch r.(type) {
	case Circle: // Assertion
		fmt.Println("It's a Circle")
	case Triangle:
		fmt.Println("It's a Triangle")
	case Square:
		fmt.Println("It's a Square")
	}
	fmt.Println(r)
	fmt.Println(r.area())
}

func main() {
	s := Square{
		size: 10.05,
	}
	display(s)

	fmt.Println(s.area())

	c := Circle{
		radius: 10.05,
	}
	display(c)

	fmt.Println(c.area())

	t := Triangle{
		breadth: 125,
		height:  60,
	}

	display(t)
	fmt.Println(t.area())
}
