// This code would fail
package main

import (
	"fmt"
)

// func scope() {
// 	var myVariable int = 10
// }

func main() {
	// fmt.Println(myVariable) // this will fail as myVariable is not in main function
	// /* rgarimel-mn3:Orielly_training rgarimel$ go run scope.go 
	//    # command-line-arguments
	//    ./scope.go:13:14: undefined: myVariable
	// */

	// Solution - declare before the block of code
	var myVariable int = 10
	fmt.Println("myVariable is ", myVariable)


	// Scope of the variable is limited to the if block alone
	// if (20 > 10) {
	// 	status := "True"
	// } else {
	// 	status := "False"	
	// }

	// fmt.Println(status) // Variable undefined, as the scope of status variable is limited to the "if" block alone.

	// Solution - declare before the block of code
	var status string
	if (20 > 10) {
		status = "True"
	} else {
		status = "False"	
	}

	fmt.Println(status) 

	// Scope of variables is limited to for loop alone

	// for x := 0; x < 10; x ++ {
	// 	var value = x
	// }
	// fmt.Println(value)
	// rgarimel-mn3:Orielly_training rgarimel$ go run scope.go 
	// # command-line-arguments
	// ./scope.go:34:14: undefined: value
	// rgarimel-mn3:Orielly_training rgarimel$ 

	// "value" variable is limited to the for loop alone


	// Solution - declare before the block of code
	var value int 
	for x := 0; x < 10; x ++ {
		value = x
	}
	fmt.Println("value = ", value)
}