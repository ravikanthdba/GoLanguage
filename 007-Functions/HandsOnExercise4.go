/*

Create a user defined struct with
the identifier “person”
the fields:
first
last
age
attach a method to type person with
the identifier “speak”
the method should have the person say their name and age
create a value of type person
call the method from the value of type person

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

func (p person) speak() string {
	return "My name is " + p.first_name + " " + p.last_name
}

func main() {

	p1 := person{
		first_name: "Ravikanth",
		last_name:  "Garimella",
		age:        32,
	}

	fmt.Println(p1)
	name := p1.speak()
	fmt.Println(name)
}
