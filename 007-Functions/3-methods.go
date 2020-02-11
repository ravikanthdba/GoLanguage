/*

	Methods

*/

package main

import (
	"fmt"
)

type person struct {
	first_name string
	last_name  string
	age        int
}

type secretAgent struct {
	person
	ltk bool
}

/*
 function format is :
 func (r receiver) <function_name> (parameters) (return types(s)) { code }

 In the below example:

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first_name, " ", s.last_name);
}

It means - for all valuees of type secretAgent, we can call the speak method

*/

func (s secretAgent) speak() {
	fmt.Println("I am ", s.first_name, " ", s.last_name)
}

func main() {

	sa1 := secretAgent{
		person: person{
			first_name: "James",
			last_name:  "Bond",
			age:        32,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()

	sa2 := secretAgent{
		person: person{
			first_name: "Miss",
			last_name:  "MoneyPenny",
			age:        27,
		},
		ltk: false,
	}

	fmt.Println(sa2)
	sa2.speak()
}
