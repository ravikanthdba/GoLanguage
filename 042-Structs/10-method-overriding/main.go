/*
	In this case, there is only one method called greeting(), but it is defined by 2 types: one by person struct and the other by agent struct.
	whenever agent object is called, it executes a different greeting(), and whenever person calls greeting(), there is a different method executed
*/

package main

import "fmt"

type person struct {
	first string
	last  string
}

type agent struct {
	person
	licenseToKill bool
}

func (p person) greeting() string {
	return "Inner type says: Howdy!!!"
}

func (a agent) greeting() string {
	return "Outer type says Hello How do u do!!!"
}

func main() {
	p1 := agent{
		person: person{
			first: "Ravikanth",
			last:  "Garimella",
		},
		licenseToKill: true,
	}

	fmt.Println(p1)
	fmt.Println(p1.greeting())

	p2 := person{
		first: "bhargavi",
		last:  "bharatula",
	}

	fmt.Println(p2)
	fmt.Println(p2.greeting())
	fmt.Println(p1.person.greeting()) // this can be overriden by p1.person => it is of type person, and its corresponding method is executed
}
