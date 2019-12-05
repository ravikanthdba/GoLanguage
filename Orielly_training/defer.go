/* defer prints the line at the end of the function. In the below example, if "defer" is not written, the order of execution is:
Goodbye
Hello
Nice Weather!!

After "defer", the order is given as:

Hello
Nice Weather!!
GoodBye

This is majorly used to close a file, after opening, so that we don't keep any open file handlers. Also, we use this for database connections, to close the connections after opening, selecting and closing.
Without this, we end up opening multiple database connections, whuch is performance impactin

*/

package main

import (
	"fmt"
)

func socialize() {
	defer fmt.Println("Goodbye")
	fmt.Println("Hello")
	fmt.Println("Nice Weather!!")
}

func main() {
	socialize()
}
