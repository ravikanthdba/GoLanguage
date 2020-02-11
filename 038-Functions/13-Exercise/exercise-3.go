package main

import (
	"fmt"
)

func greatest(variables ...int) int {
	var greatest_variable int = 0

	for _, variable := range variables {
		if variable > greatest_variable {
			greatest_variable = variable
		} else {
			continue
		}
	}

	return greatest_variable
}

func main() {
	fmt.Println(greatest(2, 3, 4))
	fmt.Println(greatest(2, 3, 1))
	fmt.Println(greatest(20, 30, 1, 90))
}
