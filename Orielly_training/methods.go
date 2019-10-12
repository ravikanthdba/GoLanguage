package main

import (
	"fmt"
)

type myType string

func (m myType) sayHello() {
	fmt.Println("hello from " + m)
}

// Methods for an embedded type also get promoted too

type Coordinates struct {
	Latitudes float64
	Longitudes float64
}

type Location struct {
	name string
	Coordinates
}

func (c Coordinates) printCoordinates() string {
	return fmt.Sprintf("(%0.2f, %0.2f)\n", c.Latitudes, c.Longitudes)
}

func main() {
	// Below code would'nt work as string1 and string2 are not of type myType

	// var string1 string
	// var string2 string

	// string1.sayHello()
	// string2.sayHello()


	value1 := myType("Ravikanth")
	value2 := myType("Bhargavi")

	value1.sayHello()
	value2.sayHello()

	// type doesn't inherit for another type, hence the bwlow commented code would not work

	type myType2 myType
	variable1 := myType2("Ravikanth")
	fmt.Println(variable1)
	fmt.Printf("%T\n", variable1)
	// variable1.sayHello()


	var l Location
	l.name = "India"
	l.Latitudes = 20.5937
	l.Longitudes = 78.9629

	fmt.Println(l.printCoordinates()) // l has "name", and the variables of the "Coordinates" struct as well. only the variables with "Coordinate" struct of l, get passed to the function.

	// Exercise: https://play.golang.org/p/uo1igDqgBv7
}