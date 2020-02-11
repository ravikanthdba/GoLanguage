/* Every rune code of the upper case alphabet should be printed 3 times */

package main

import (
	"fmt"
)

func main() {

	for i := 65; i <= 90; i++ {
		fmt.Printf("%d: \n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%U\t'%c'\n", i, i)
		}
	}
}
