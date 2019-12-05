package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) helloworld() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{first: "Ravikanth", last: "Garimella", age: 32}
	p2 := person{first: "Bhargavi", last: "Bharatula", age: 28}

	fmt.Println(p1.helloworld())
	fmt.Println(p2.helloworld())
}
