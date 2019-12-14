package _1_example_1

import (
	"fmt"
	"errors"
	"log"
	"math"
)

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Fatal(err)
	}
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, &norgateMathError{
			Latitude:  "12345",
			Longitude: "54321",
			err:       errors.New("This is an error, as square root of negative number is not possible"),
		}
	}
	return math.Sqrt(n), nil
}


type norgateMathError struct {
	Latitude, Longitude string
	err error
}

func (n *norgateMathError) Error() string {
	return fmt.Sprintf("Error returned with Latitude: %s, Longitude: %s and message: %v\n", n.Latitude, n.Longitude, n.err)
}



