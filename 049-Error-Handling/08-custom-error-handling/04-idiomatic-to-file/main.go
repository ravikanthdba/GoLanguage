package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math"
	"os"
)

var logFileName string = "logging.txt"
var errMessage = errors.New("Square root of a negative number is not formed")

func main() {
	logfile, err := os.Create(logFileName)
	if err != nil {
		fmt.Println("Unable to open log file")
		return
	}
	fmt.Println("Opened the logfile: ", logFileName)
	log.SetOutput(logfile)

	sqrtNumber, err := sqrt(-10)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The square root of %f is %f", -10, sqrtNumber)
}

func sqrt(number float64) (float64, error) {
	if number < 10 {
		return 0, errMessage
	}
	return math.Sqrt(number), nil
}
