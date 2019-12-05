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
	p1 := &person{first: "Ravikanth", last: "garimella", age: 32}
	fmt.Println(p1)
	fmt.Println(*p1)

	fmt.Printf("The type of p1 is %T\n", p1)
	fmt.Printf("The type of *p1 is %T\n", *p1)

	fmt.Println(p1.first, p1.last, p1.age) // pulls data directly out of pointer address (memory). Internally does *p1.first
}
