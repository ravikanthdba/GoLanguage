/*

Create a value and assign it to a variable.
Print the address of that value.

*/

package main

import (
	"fmt"
)

func main() {

	var x int

	x = 33
	fmt.Printf("The value of x is %d\n", x)
	fmt.Printf("The address of x is %#x\n", &x)
	fmt.Printf("The address of x is %#X\n", &x)
}
