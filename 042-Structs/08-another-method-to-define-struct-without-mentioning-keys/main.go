/* The structs can be defined without even mentioning the key, but it is not recommended */

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
	p1 := person{"ravikanth", "garimella", 32}
	p2 := person{"bhargavi", "bharatula", 28}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
