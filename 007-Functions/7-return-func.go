/*

   In several cases, a function can even return a function.
   The syntax is as follows:

   func (r receiver) <function_name> (parameters) func() <datatype of what the func returns> {
       code
   }


   Example is as follows:

*/

package main

import (
	"fmt"
)

func main() {

	// function1 := foo();
	// fmt.Println("The output of function1 is ", function1);

	// function2 := bar()
	// fmt.Println("The output of function2 is ", function2);
	// fmt.Println("The output of function2 upon executing is ", function2());

	fmt.Println("\n\n")
	fmt.Println("cleaning the code:")

	fmt.Println("The output of cleaning the code is to just reduce using assignment of another variable.")
	fmt.Println(bar()())

}

// func foo() int {
//     return 48;
// }

func bar() func() string {

	returnfunction := func() string {
		return "This is through function return"
	}

	return returnfunction

}
