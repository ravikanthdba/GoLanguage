package main

import (
	"fmt"
)

func main() {
	var friend string = "hello1"

	switch {
	case len(friend) == 2:
		fmt.Println("The length of my friends name is 2")

	case len(friend) > 2 && len(friend) < 5:
		fmt.Println("The length of my friends is greater than 2 but less than 5")

	case friend == "hello1", friend == "hello2":
		fmt.Println("Hi Hello1 or Hello2")

	default:
		fmt.Println("Greater than 5")
	}
}
