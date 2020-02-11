/*
   The below example shows interfaces and polymorphism

*/

package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

/* In the below code, 2 different types have same functino - this is called Polymorphism
   In this case, person and secretAgeent have function speak()

*/

func (s secretAgent) speak() {
	fmt.Println("I am", s.first, s.last)
}

func (p person) speak() {
	fmt.Println("Person: ", p.first, p.last)
}

// The below code says, any variable of a type having a speak() function is also of type human
/* When we have an empty interface, every other type implements the empty interface
    In thee below code, we had commented the speak method, which means, every type is considered as human, irrespectiive of what methods he has.

    // type human interface {
	//     // speak();
	// }

    So even if you do not define a method for person, as above, this would still execute.

*/

type human interface {
	speak()
}

// In the below code, any variable of a type human, has access to the bar variable
//In the below code: h.(person).first is called assertion. Which means, I know h is of type person, so print the element first of the type person

func bar(h human) {
	switch h.(type) {
	case person:
		fmt.Println("I am in human now", h.(person).first)
	case secretAgent:
		fmt.Println("I am in human now", h.(secretAgent).first)
	}
}

func main() {

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	fmt.Println(sa1)
	sa1.speak()
	bar(sa1)

	fmt.Println("\n\n")

	sa2 := secretAgent{
		person: person{
			first: "Miss",
			last:  "MoneyPenny",
			age:   27,
		},
		ltk: false,
	}

	fmt.Println(sa2)
	sa2.speak()
	bar(sa2)

	/* If I wouldn't have implemented speak on person type, then this code would fail. As bar function is only for type human.
	   Human is an interface, which is created with variables having speak() function
	*/

	fmt.Println("\n\n")
	p1 := person{
		first: "Ravikanth",
		last:  "Garimella",
		age:   32,
	}
	fmt.Println(p1)
	bar(p1)

	// conversion

	type hotdog int
	var x hotdog = 42
	fmt.Println("the value of x is ", x)
	fmt.Printf("the type of x is %T\n", x)

	var y int
	y = int(x)
	fmt.Println("the value of y is ", y)
	fmt.Printf("the type of y is %T\n", y)
}
