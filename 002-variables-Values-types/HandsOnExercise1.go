/*
   Hands On Exercise1 - USing short declaration operator, assign values to variables x, y, and z
*/

package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	// Print variables using single print statement
	// Print variables using multiple print statement

	fmt.Println("Single Line Statement")
	fmt.Println("---------------------")
	fmt.Printf("The value of x is : %d, y is %s and z is %t\n\n", x, y, z)
	fmt.Println("Multi Line Statement")
	fmt.Println("---------------------")
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
