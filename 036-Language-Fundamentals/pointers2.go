package main

import (
	"fmt"
)

func main() {
	var a int = 42
	fmt.Println("The value of a is: ", a)
	fmt.Println("The address of a is: ", &a)

	fmt.Println("\n")

	var b *int = &a
	fmt.Println("The value of b is: ", b)
	fmt.Println("The address of b is: ", &b)
	fmt.Println("The value of *b is: ", *b)
	fmt.Println("\n")

	fmt.Println("Changing the value of b")

	fmt.Println("\n")

	*b = 99
	fmt.Println("The value of b is: ", b)
	fmt.Println("The address of b is: ", &b)
	fmt.Println("The value of *b is: ", *b)

	fmt.Println("\n")

	fmt.Println("The value of a is: ", a)
	fmt.Println("The address of a is: ", &a)

}
