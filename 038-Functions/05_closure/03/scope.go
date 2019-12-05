package main

import (
	"fmt"
)

func main() {
	var x int

	increment := func() int {
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())
}

/*
   closure helps reduce the scope of a variable. The more the scope of the variable is reduced, the better. In the previous example, x is package level scoped.
   In this example, x is now limited to main function.

*/
