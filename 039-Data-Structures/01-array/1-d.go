package main

import (
	"fmt"
)

func main() {
	var x = [5]int{1, 2, 3, 4, 5}
	fmt.Println(len(x))
	for _, variable := range x {
		fmt.Println(variable)
	}
}
