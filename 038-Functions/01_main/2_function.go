package main

import (
	"fmt"
)

func main() {
	greet("Ravikanth")
	greet("Viraj")
	greet("Bhargavi")
}

func greet(name string) {
	fmt.Println("Welcome ", name)
}
