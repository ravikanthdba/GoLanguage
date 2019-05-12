/*

Create a func which returns a func
assign the returned func to a variable
call the returned func

*/

package main

import (
	"fmt"
)

func main() {

	var x, y int;
	x = 20; y = 30;
	fmt.Println("The contents of the function is :", hello()());
	fmt.Printf("The sum of %d and %d is %d\n", x, y, sum(x, y)());

}

func hello() func() string {
	return func() string {
		return "My name is Ravi, this is my first golang program.."
	}
}


func sum(x, y int) func() int {
	return func() int {
		return (x + y);
	}
}