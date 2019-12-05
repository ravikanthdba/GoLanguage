package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Println("Enter the number:")
	fmt.Scanf("%d", &number)
	if (number % 4) == 0 {
		fmt.Println("Linkedin")
	} else if number%4 == 1 {
		fmt.Println("Linked")
	} else {
		fmt.Println("In")
	}
}
