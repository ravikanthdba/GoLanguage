package main

import (
	"fmt"
)

func main() {
	var x = [5]int{1, 2, 3, 4, 5}
	for _, variable := range x {
		fmt.Printf("The variable %d is at memory location %d\n", variable, &variable)
	}
}
