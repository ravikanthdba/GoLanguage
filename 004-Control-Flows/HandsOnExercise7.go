/* Write a program with if else if */

package main

import (
	"fmt"
)

func main() {
	var integer int;

	fmt.Println("Enter the integer: ")
	fmt.Scanf("%d", &integer);


	if (integer % 4 == 0 && integer % 6 == 0) {
		fmt.Println("Linkedin");
	} else if (integer % 4 == 0) {
		fmt.Println("Linked");
	} else if (integer % 6 == 0) {
		fmt.Println("In");
	} else {
		fmt.Println("Invalid number");
	}
}