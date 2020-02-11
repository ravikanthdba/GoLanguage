package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

var (
	errorMessageDivideByZero = errors.New("You are trying to divide by zero here, we cannot divide by zero")
)

func main() {
	var (
		a float64 = 2050
		b float64 = 0.000
	)

	divisioning, err := division(a, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%f / %f = %f\n", a, b, divisioning)
}

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errorMessageDivideByZero
	}
	return a / b, nil
}
