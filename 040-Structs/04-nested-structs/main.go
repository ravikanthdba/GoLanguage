package main

import (
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

type telugu struct {
	person
	First  string
	Second string
}

type kannada struct {
	person
	First  string
	Second string
}

func main() {
	p1 := telugu{
		person: person{
			First: "Ravikanth",
			Last:  "Garimella",
			Age:   32,
		},
		First:  "telugu",
		Second: "hindi",
	}

	p2 := kannada{
		person: person{
			First: "xyz",
			Last:  "abc",
			Age:   33,
		},
		First:  "kannada",
		Second: "hindi",
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.First, p1.Second, p1.person.First, p1.person.Last, p1.person.Age)
	fmt.Println(p2.First, p2.Second, p2.person.First, p2.person.Last, p2.person.Age)

	/* p1.First will give you the value of First from p1 struct, p2.First will give you the value of First from p2 struct,
	but p1.person.First - would give you the data from the person struct
	This is called promotion
	*/
}
