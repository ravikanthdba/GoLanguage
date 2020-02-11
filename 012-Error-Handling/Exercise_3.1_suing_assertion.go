package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: ", ce.info)
}

func main() {
	c1 := customErr{
		info: "Need some more coffee",
	}

	fmt.Printf("The type of c1 is %T\n", c1)

	foo(c1)
}

func foo(e error) {
	fmt.Printf("The type of e is %T\n", e)
	fmt.Println("From the function foo(), using assertion the error is now: ", e.(customErr).info)
	// e will not have the Info variable, we need to assert that e is of type Customerr, which has the info variable.
}
