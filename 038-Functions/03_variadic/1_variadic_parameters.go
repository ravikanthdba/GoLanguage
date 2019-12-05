/*
   1) denoted by ...
   2) pass unlimited values to the function
*/

package main

import (
	"fmt"
)

func average(listVariables ...float64) float64 { // accepts variadic parameters
	var total float64
	var count int

	for _, variables := range listVariables {
		total += float64(variables)
		count += 1
	}

	return total / float64(count)
}

func main() {
	n := average(100, 200, 150, 250, 300)
	fmt.Println(n)

	n = average(10, 20, 30)
	fmt.Println(n)
}
