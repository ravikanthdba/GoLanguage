package main

import (
	"fmt"
)

func printStatement(friend string) {
	fmt.Println("you've enteered the option as: ", friend)
	switch friend {
	case "Ravi":
		fmt.Println("It's Ravikanth")
	case "Medi":
		fmt.Println("You've entered Medi")
	case "Jenny":
		fmt.Println("What's up Jenny")
	default:
		fmt.Println("Do you have Friends")
	}
}

func main() {
	printStatement("Ravi")
	printStatement("ravi")
	printStatement("Medi")
	printStatement("Jenny")
}
