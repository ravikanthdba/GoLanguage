package main

import (
	"fmt"
	"log"
)

func main() {
	square_root, err := sqrt(-10)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(square_root)
}

func sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("number is less than 0, and hence the square root is not possible. Number passed is: %f", number)
	}
	return 42, nil
}
