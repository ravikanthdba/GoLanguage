/*

Build and use an anonymous func

*/

package main

import (
	"fmt"
)

func main() {

	var x int = 50

	func(variable int) {
		fmt.Println("The value of x is ", variable)
	}(x)

	func() {
		fmt.Println("This is an anonymous function")
	}()

}
