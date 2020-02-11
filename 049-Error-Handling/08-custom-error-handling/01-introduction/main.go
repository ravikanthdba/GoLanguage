/* error type in go are just like any other types in Go. We can create custom errors by using "errors" package in Go */

package main

import (
	"github.com/pkg/errors"
	"log"
	"math"
)

func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("The number is a negative number, hence square root cannot be found")
	}
	return math.Sqrt(n), nil
}
