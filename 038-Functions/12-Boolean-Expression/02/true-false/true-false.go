package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("This executes")
	}

	if false {
		fmt.Println("This doesn't execute")
	}

	var _ int = 14
	if true {
		fmt.Println("X value is assigned..")
	}

	var _ int = 0
	if true {
		fmt.Println("X value is assigned..")
	}
}
