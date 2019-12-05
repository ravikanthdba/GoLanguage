/*


Using a COMPOSITE LITERAL:
create a SLICE of TYPE int
assign 10 VALUES
Range over the slice and print the values out.
Using format printing
print out the TYPE of the slice

*/

package main

import (
	"fmt"
)

func main() {

	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	for value := range slice {
		fmt.Println(slice[value])
	}

	fmt.Printf("The type of slice is %T\n", slice)
}
