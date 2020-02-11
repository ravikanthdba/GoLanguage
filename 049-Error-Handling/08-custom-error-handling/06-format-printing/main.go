	package main

import (
	"fmt"
	"log"
)

func main() {
	var (
		a float64 = 20
		b float64 = 0
	)

	value, err := division(a, b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("denominator is 0 now, and %f / %f is not possible now\n", a, b)
	}
	return a / b, nil
}
