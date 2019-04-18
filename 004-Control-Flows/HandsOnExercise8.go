/* Program with Switch statement with no condition */

package main

import (
	"fmt"
)

func main() {

	switch {
		case true:
			fmt.Println("Printing it!!");
		
		case false:
			fmt.Println("Not Printing it!!");

		default:
			fmt.Println("Default condition!!");
	}
}