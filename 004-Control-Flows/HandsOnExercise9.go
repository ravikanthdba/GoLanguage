/* Program with a seitch expression with a condition */

package main

import (
	"fmt"
)

func main() {

	var favSport string
	fmt.Println("Enter your favourite sport:")
	fmt.Scanf("%s", &favSport)
	switch favSport {
	case "Cricket":
		fmt.Println("You've entered cricket")

	case "Football":
		fmt.Println("You've entered Football")

	case "Hockey":
		fmt.Println("You've entered Hockey")

	default:
		fmt.Println("Invalid option, you were supposed to choose from 'Cricket', 'Hockey', 'Football'")
	}
}
