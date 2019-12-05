/*

Assign a func to a variable, then call that func

*/

package main

import (
	"fmt"
)

func main() {

	x := func() {
		fmt.Println("This is an anonymous function assiigned to a variable x")
	}

	x()
}
