package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{first: "Ravi", last: "Garimella", age: 32}
	p2 := person{first: "Swarna Latha", last: "Garimella", age: 53}
	p3 := person{first: "Nagabhushanam", last: "Garimella", age: 63}
	p4 := person{first: "Bhargavi", last: "Bharatula", age: 28}
	p5 := person{first: "Viraj", last: "Bharatula", age: 1}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(p3.first, p3.last, p3.age)
	fmt.Println(p4.first, p4.last, p4.age)
	fmt.Println(p5.first, p5.last, p5.age)
}

type Hello struct {
	World string
	Name string
}