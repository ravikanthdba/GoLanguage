package main

import (
	"fmt"
)

func main() {
	if true && true {
		fmt.Println("This prints..")
	}

	if true && false {
		fmt.Println("This doesn't prints..")
	}

	if false && true {
		fmt.Println("This doesn't prints..")
	}

	if false && false {
		fmt.Println("This doesn't prints..")
	}

	fmt.Println("\n")
	fmt.Println("Truth Table:")
	fmt.Println("=============")
	fmt.Println("true and true   = \t", true && true)
	fmt.Println("true and false  = \t", true && false)
	fmt.Println("false and true  = \t", false && true)
	fmt.Println("false and false = \t", false && false)
}
