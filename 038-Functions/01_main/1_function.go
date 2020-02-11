/*
main function is an entry point to the program.
There cannot be more than one main function.
There can be code without the maun function, those are mostly packages. Those are just compiled and a binary is generated.

Example syntax of a function:

func (<optional> receiver receiver_type) function_name(<optional> arguments list) (<optional> return_type1, return_type2..) {
    code block
}


*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Functions!!")
}
