package main

import "fmt"

type gator int
type flamingo bool

func (g gator) greeting() {
	fmt.Println("Hello I'm a gator")
}

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
	greeting()
}

func bayou(s swampCreature) {
	switch s.(type) {
	case gator:
		fmt.Println("This is a gator, which is passed")
	case flamingo:
		fmt.Println("This is a flamingo, which is passed")
	}

	s.greeting()
}

func main() {
	var g1 gator = 10
	fmt.Printf("The type of g1 is %T and value of g1 is %d\n", g1, g1)

	var g2 gator = 100
	fmt.Printf("The type of g2 is %T and the value of g2 is %d\n", g2, g2)

	var x int = 2000
	fmt.Printf("The type of x is %T and the value of x is %d\n", x, x)

	/* This code would not work as the variable is int and the value you are assigning is gator, and since the type is different, this would not work

	# command-line-arguments
	./11_types.go:19:4: cannot use g1 (type gator) as type int in assignment
	*/
	//x = g1
	//fmt.Printf("The type of x is %T and the value of x is %d\n", x, x)

	// This would work, as we are converting variable of gator type to int
	x = int(g1)
	fmt.Printf("The type of x is %T and the value of x is %d\n", x, x)

	g1.greeting()
	g2.greeting()

	var f1 flamingo = true
	var f2 flamingo = false

	bayou(g1)
	bayou(g2)
	bayou(f1)
	bayou(f2)
}
