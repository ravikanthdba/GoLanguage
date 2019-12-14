package main

import "fmt"

/* person is an inner type struct */
type person struct {
	first string
	last  string
}

/* agent is an outer type struct */
/* outer type overrides the inner type*/
type agent struct {
	person
	licenseToKill bool
}

func main() {
	p1 := agent{
		person: person{
			first: "Ravi",
			last:  "garimella",
		},
		licenseToKill: true,
	}

	fmt.Println(p1)

	/* alternate method without defining the keys */

	p2 := agent{
		person{
			"Ravi",
			"Garimella",
		},
		true,
	}

	fmt.Println(p2)
}
