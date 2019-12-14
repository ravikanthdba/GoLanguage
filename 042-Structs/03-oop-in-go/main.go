/*
	Go is an object oriented programming language, it has :
	- encapsulation - means state and behaviour - has exported / unexported variables - means state and behaviour
	- polymorphism - one function implements multiple types
	- reusability
	- overriding (promotion)
*/

package main

import (
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int
}

type Agent struct { // outer type
	Person // Anonymous, promoted, inner type. Whenever an anonymous struct is added, all the inner types are promoted to the outer struct
	// Any object created in Agent type, can access all the variables with Person type
	// Person is inherited in Agent
	LicenseToKill bool
}

func main() {
	p1 := Agent{
		Person: Person{
			First: "Ravikanth",
			Last:  "Garimella",
			Age:   32,
		},
		LicenseToKill: true,
	}

	p2 := Agent{
		Person: Person{
			First: "Swarna Latha",
			Last:  "Garimella",
			Age:   53,
		},
		LicenseToKill: true,
	}

	p3 := Agent{
		Person: Person{
			First: "Nagabhushanam",
			Last:  "Garimella",
			Age:   63,
		},
		LicenseToKill: false,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)
}
