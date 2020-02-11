package main

import "fmt"

type normalPerson struct {
	firstname string
	lastname  string
	age       int
}

type secretiveAgent struct {
	normalPerson
	licenseToKill bool
}

func (n normalPerson) speak() {
	fmt.Println("Normal Person: ", n.firstname+" "+n.lastname+" says: Good morning!!!")
}

func (s secretiveAgent) speak() {
	fmt.Println("Secret Agent: ", s.lastname+" "+s.firstname+" says: I'm a secret agent!!")
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	n1 := normalPerson{
		firstname: "Ravikanth",
		lastname:  "Garimella",
		age:       32,
	}

	sa1 := secretiveAgent{
		normalPerson: normalPerson{
			firstname: "Viraj",
			lastname:  "Nandan",
			age:       1,
		},
		licenseToKill: true,
	}

	saySomething(n1)
	saySomething(sa1)
}
