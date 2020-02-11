/*

defer keyword - defers execution of a function
For example, if we open a file and we forget to close it, then it wastes too much memory.
In Go, this is a very good use case, where you can define closing the file immediately after openin g it and usage of "defer keyword", will make it execute when the function is about to end.

Below is the example

*/

package main

import (
	"fmt"
)

func main() {

	fmt.Println("Without Defer keyword, it executes in sequential order")
	foo()
	bar()

	fmt.Println("\n\n")
	fmt.Println("Now defer is applied on foo function")

	defer foo()
	bar()

	fmt.Println("\n\n")
	fmt.Println("In the above example, after application of defer on foo function, foo is executed before the function main ends, and hence bar is executed before foo")
}

func foo() {
	fmt.Println("The text is displayed from foo function")
}

func bar() {
	fmt.Println("The text is displayed from bar function")
}
