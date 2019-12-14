package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	var number float64
	fmt.Printf("Enter the number you eant to find the Square root for: ")
	fmt.Scanf("%f", &number)

	num, err := sqrt(number)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The square root of the number: %f is %f\n", number, num)
}

func sqrt(number float64) (float64, error) {
	if number < 0 {
		norgateErrorMessage := fmt.Errorf("Cannot get the square root of negative number %v\n", number)
		return 0, norgateErrorMessage
	}
	return math.Sqrt(number), nil
}
