package main

import (
	"fmt"
)

func main() {
	data := []float64{22, 20, 30, 40, 12}
	avg := average(data...) // variadic arguments
	fmt.Println(avg)
}

func average(sf ...float64) float64 {
	var total float64

	for _, variable := range sf {
		total += variable
	}

	return total / float64(len(sf))
}
