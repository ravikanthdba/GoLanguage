package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math"
)

var ErrMessage = errors.New("Negative integer, there is no square root of a negative number")

func main() {
	sqrtNumber, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Square Root is: ", sqrtNumber)
}

func sqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, ErrMessage
	}
	return math.Sqrt(number), nil
}
