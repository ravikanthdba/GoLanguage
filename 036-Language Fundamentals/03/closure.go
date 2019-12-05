package main

import (
	"fmt"
)

// function wrapper returns a function. In Go, just like we return an int type, float type, string type, we can also return a function, as function is also a type
func wrapper() func() int {
	x := 0              // Assign a value 0 to x
	return func() int { // returns a function
		x++
		return x
	}
}

func main() {
	increment := wrapper() // assigns wrapper function() to increment variable.
	fmt.Println("The value of increment variable is: ", increment)
	fmt.Printf("The type of increment variable is: %T\n", increment)
	fmt.Println(increment())
	fmt.Println(increment())
}
