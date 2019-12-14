package main

import (
	"fmt"
	"github.com/pkg/errors"
	"math"
)

func main() {
	var (
		a float64 = 10
		b float64

		n float64 = -10.23
	)

	_, divisionerr := division(a, b)
	if divisionerr != nil {
		fmt.Println(divisionerr)
	}

	_, sqrterr := sqrt(n)
	if sqrterr != nil {
		fmt.Println(sqrterr)
	}

}

type mathError struct {
	function string
	err error
}

func (d *mathError) Error() string {
	return fmt.Sprintf("From the function %v, the error is: %v", d.function, d.err)
}

func division(a, b float64) (float64, error) {
	if b == 0 {
		return 0, &mathError{
			function: "division",
			err:      errors.New("Trying to divide by 0?? are you crazy??"),
		}
	}
	return (a / b), nil
}


func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, &mathError{
			function: "sqrt",
			err:      errors.New("Trying to find the square root of negative number??, are you crazy"),
		}
	}
	return math.Sqrt(n), nil
}
