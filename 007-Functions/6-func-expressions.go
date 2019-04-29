/*

	We can assign a function to a variable, like we can assign an integer to a variable (or) a string to a variable

*/

package main

import (
	"fmt"
)

func main() {

	f := func() {
		fmt.Println("This is my first func expression")
	}
	f()
	fmt.Printf("Func expression is of type: %T\n", f)
	fmt.Printf("The value of f is : %d \n", f)  // This prints the memory address of f
	fmt.Printf("The value of f is : %#x \n", f) // This prints the memory address of g

	fmt.Println("\n\n")

	g := func(x int) {
		fmt.Printf("The value of x is %d\n", x)
	}
	g(20)
	fmt.Printf("Func expression is of type: %T\n", g)
	fmt.Printf("The value of g is : %d \n", g)  // This prints the memory address of g
	fmt.Printf("The value of g is : %#x \n", g) // This prints the memory address of g

}
