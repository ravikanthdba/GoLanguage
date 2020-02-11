/*

Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:

"James", "Bond", "Shaken, not stirred"
"Miss", "Moneypenny", "Helloooooo, James."

Range over the records, then range over the data in each record.


*/

package main

import (
	"fmt"
)

func main() {

	multidimensionalslice1 := []string{"James", "Bond", "Shaken, not stirred"}
	multidimensionalslice2 := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	multidimensionalslice := [][]string{multidimensionalslice1, multidimensionalslice2}

	for i := 0; i < len(multidimensionalslice); i++ {
		for j := 0; j <= len(multidimensionalslice); j++ {
			fmt.Println(multidimensionalslice[i][j])
		}
	}

}
