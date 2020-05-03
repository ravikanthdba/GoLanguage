/* unexpected errors are handled by Panic. It is used by built in "panic" function.	Immediately contact the administrator
Anything after panic would not run
setting "defer", still runs even despite the panic

Numerical value shown below, shows the steps of execution
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")                       // 1
	defer fmt.Println("Not with this airline") // 3
	defer fmt.Println("Defering again")        // 2
	panic("I need to get out of the party.")   // 4
	fmt.Println("Goodbye!!")                   // 5
}
