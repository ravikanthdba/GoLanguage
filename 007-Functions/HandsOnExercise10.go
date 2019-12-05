/*

Closure is when we have “enclosed” the scope of a variable in some code block.
For this hands-on exercise, create a func which “encloses” the scope of a variable:

*/

package main

import (
	"fmt"
)

func main() {
	x := 42
	fmt.Println("From main x is:", x)
	{
		fmt.Println("From block x is:", x)
		y := 100
		fmt.Println("From block y is:", y)
	}
	// fmt.Println("From main y is:", y); // Fails because y is not defined in main function
}
