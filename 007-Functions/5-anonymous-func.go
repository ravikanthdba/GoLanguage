/*

Anonymous function - is a function without a name

In the below example, main function has a function and it is without a name

*/

package main

import (
	"fmt"
)

func main() {

	var x int = 42
	foo(x)

	func() {
		fmt.Println("This is an Anonymous function without parameters and arguments")
	}()

	func(variable int) {
		fmt.Printf("Printing variable: %d from the Anonymous function\n", variable)
	}(50)

}

func foo(variable int) {
	fmt.Printf("Printing variable: %d from the function foo\n", variable)
}
