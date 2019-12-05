package main

import (
	"fmt"
)

func main() {
	fmt.Println(greet("Ravikanth Garimella"))
	fmt.Println(repeat("Ravikanth Garimella", 5))
}

func greet(name string) string {
	return "Welcome to the Function " + name
}

func repeat(name string, number int) string {
	var output string
	for i := 0; i < number; i++ {
		output += name
	}
	return output
}
