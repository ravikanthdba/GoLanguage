package main

import (
	"fmt"
)

func main() {
	numbers := []float64{10, 20, 30, 40, 50}
	fmt.Println(average(numbers))
}

func average(sf []float64) float64 {
	var total float64

	for _, variable := range sf {
		total += variable
	}

	return total / float64(len(sf))
}
