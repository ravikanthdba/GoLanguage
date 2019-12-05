package main

import (
	"fmt"
)

func main() {
	if true || true {
		fmt.Println("This prints..")
	}

	if true || false {
		fmt.Println("This prints..")
	}

	if false || true {
		fmt.Println("This prints..")
	}

	if false || false {
		fmt.Println("This prints..")
	}

	fmt.Println("\n")
	fmt.Println("Truth Table:")
	fmt.Println("=============")
	fmt.Println("true or true   = \t", true || true)
	fmt.Println("true or false  = \t", true || false)
	fmt.Println("false or true  = \t", false || true)
	fmt.Println("false or false = \t", false || false)
}
