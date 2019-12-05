package main

import (
	"fmt"
)

func difference(variable1, variable2 int) int {

	if variable1 > variable2 {
		return (variable1 - variable2)
	} else {
		return (variable2 - variable1)
	}
}

func main() {
	var variable1 int = 987
	var variable2 int = 564

	fmt.Println(difference(variable1, variable2))
}
