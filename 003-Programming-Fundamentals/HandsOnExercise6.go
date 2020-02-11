/* Using IOTA create constants for the last 4 years and print the constant values */

package main

import (
	"fmt"
)

const (
	_      = iota
	year_0 = 2013
	year_1 = 2013 + iota
	year_2 = 2013 + iota
	year_3 = 2013 + iota
)

func main() {
	fmt.Println("Year_0 is ", year_0)
	fmt.Println("Year_1 is ", year_1)
	fmt.Println("Year_2 is ", year_2)
	fmt.Println("Year_3 is ", year_3)
}
