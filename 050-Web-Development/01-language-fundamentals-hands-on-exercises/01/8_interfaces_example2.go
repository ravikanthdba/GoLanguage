package main

import (
	"fmt"
	"strconv"
)

type people struct {
	firstname string
	lastname  string
	age       int
}

type secretAgent struct {
	people
	licenseToKill bool
}

func (p people) pspeak() {
	fmt.Println(p.firstname + " " + p.lastname + " of age: " + strconv.Itoa(p.age) + " says hello!!")
}

func (sa secretAgent) saspeak() {
	fmt.Println(sa.firstname + " " + sa.lastname + " of age: " + strconv.Itoa(sa.age) + " is a secret Agent!!! Beware")
}

func main() {
	p1 := people{
		firstname: "Ravikanth",
		lastname:  "Garimella",
		age:       32,
	}

	p1.pspeak()

	p2 := secretAgent{
		people: people{
			firstname: "Viraj",
			lastname:  "Nandan",
			age:       1,
		},
		licenseToKill: false,
	}

	p2.saspeak()
}
