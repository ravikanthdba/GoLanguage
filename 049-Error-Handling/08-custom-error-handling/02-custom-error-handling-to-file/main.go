package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"math"
	"os"
)

func main() {
	logfile, err := os.Create("logger.log")
	if err != nil {
		fmt.Println("Error creating the file: ", err)
		return
	}
	fmt.Println("Logfile: ", logfile, " created.")
	log.SetOutput(logfile)


	sqrtNumber, err := sqrt(-10)
	if err != nil {
		log.Fatalf("%s\n", err)
	}

	fmt.Printf("The square of %f is %f\n", -10, sqrtNumber)
}


func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, errors.New("n is a negative number hence no square root")
	}
	return math.Sqrt(n), nil
}
