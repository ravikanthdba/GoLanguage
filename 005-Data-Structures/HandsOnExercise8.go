/*

Create a map with a key of TYPE string which is a person’s “last_first” name, and a value of TYPE []string which stores their favorite things. Store three records in your map. Print out all of the values, along with their index position in the slice.

	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
`no_dr`, `Being evil`, `Ice cream`, `Sunsets`



*/

package main

import (
	"fmt"
)

func main() {

	m := map[string][3]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}

	for mapkeys := range m {
		for favourite := 0; favourite < len(m[mapkeys]); favourite++ {
			fmt.Printf("At index position %d for the map %s, the favourite thing is %s\n", favourite, mapkeys, m[mapkeys][favourite])
		}
	}
}
