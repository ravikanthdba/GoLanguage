/*

Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop


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
		`ravikanth`:       {`Watching movies`, `programming`, `reading`},
	}

	// Delete a record

	delete(m, `no_dr`)

	fmt.Printf("%s           -  %s\n", "name", "favourite")
	for mapkeys := range m {
		for favourite := 0; favourite < len(m[mapkeys]); favourite++ {
			fmt.Printf("%s           -  %s\n", mapkeys, m[mapkeys][favourite])
		}
	}
}
