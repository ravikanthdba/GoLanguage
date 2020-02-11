package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func sayHello(p person) {
	fmt.Println(p.lastname + " " + p.firstname + " says hello!!!")
}

func main() {
	p1 := person{
		firstname: "Ravikanth",
		lastname:  "Garimella",
	}

	sayHello(p1)
}
