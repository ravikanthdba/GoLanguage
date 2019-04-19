/* Create a variable with Raw String literal and print it 

`` - If we put the String in back ticks as shown below, then you will be able to put the string in new lines.
"" - Doesn't allow that

*/

package main


import (
	"fmt"
)


func main() {
	x := `Hello World
	My name is Ravikanth Garimella
	I am learning the Go PRogramming Language`;

	fmt.Println(x);
}