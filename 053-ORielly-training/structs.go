package main

import (
	"fmt"
)

type bucket struct {
	number  int
	decimal float64
	toggle  bool
	name    string
}

// Embedded structs - struct within a struct

type Coordinates struct {
	Latitudes  float64
	Longitudes float64
}

type Landmark struct {
	Name string
	/*
		An Anonymous field, and has no name. It's just a type
	*/
	Coordinates
}

func main() {
	// Include all type of variables into one bucket

	var buck bucket

	buck.number = 64
	buck.decimal = 64.1234567890
	buck.toggle = true
	buck.name = "Ravikanth"

	fmt.Println(buck)

	var l Landmark
	l.Name = "The GooglePlex"
	l.Latitudes = 34.1234
	l.Longitudes = -122.08

	fmt.Println(l)
	fmt.Println(l.Longitudes, l.Latitudes, l.Name)

	// Exercise: https://play.golang.org/p/ugMBtoilOPI
}
