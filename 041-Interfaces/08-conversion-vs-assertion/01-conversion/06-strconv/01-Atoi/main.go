/* Atoi - takes a string and converts into integer */

package main

import (
	"fmt"
	"strconv"
)

func main() {
	integer, _ := strconv.Atoi("10")
	fmt.Printf("The type of integer is %T\n", integer)
	fmt.Printf("The value of integer is %d\n", integer)
}
