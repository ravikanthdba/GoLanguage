package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{First: "Ravikanth", Last: "Garimella", Age: 32}
	p2 := person{First: "Bhargavi", Last: "Bharatula", Age: 28}

	var p3 person /* This is a zero type
	If the datatype is int - then value = 0
	If the datatype is string - then value is ""
	If the datatype is bool - then value is false

	We use var to only declare zero type. If we need to declare with values, then we give short hand declaration. i.e. :=
	*/

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
