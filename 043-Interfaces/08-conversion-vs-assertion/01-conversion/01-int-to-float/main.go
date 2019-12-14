/* Widening conversion */

package main

import (
	"fmt"
)

func main() {
	var x int = 10
	var y float64 = 12.12345678
	fmt.Printf("%f + %f = %f\n", float64(x), y, float64(x)+y)
}
