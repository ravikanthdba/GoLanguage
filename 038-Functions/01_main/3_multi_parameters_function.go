package main

import (
	"fmt"
)

func main() {
	greet("Ravikanth", "GArimella", 32)
}

func greet(first, last string, age int) {
	fmt.Println("Hello!!!!", first, last, " of age ", age)
}
