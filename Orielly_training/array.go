// Contiguous data store

/* Limitations
- Cannot increase

*/

package main

import (
	"fmt"
)

func main() {
	var primes [3]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(primes[1])
	fmt.Println(primes)

	fmt.Println("\n")

	var booleans [2]bool
	booleans[0] = true
	booleans[1] = false
	fmt.Println(booleans)

	fmt.Println("\n")

	var strings [5]string
	strings[0] = "Ravikanth"
	strings[2] = "Garimella"
	fmt.Println(strings)
	//%#v - prints values in debug format - this is called array literal
	fmt.Printf("%#v\n", strings)

	fmt.Println("\n")

	// Default values are 0 values
	var integers [3]int
	var names [5]string
	var booleanobjects [2]bool
	fmt.Printf("%#v\n", integers)
	fmt.Printf("%#v\n", names)
	fmt.Printf("%#v\n", booleanobjects)

	fmt.Println("\n")

	// Create array using array literals
	refridgerators := [2]string{"Videocon", "LG"}
	fmt.Println(refridgerators)
	fmt.Println(refridgerators[0])
	fmt.Println(refridgerators[1])

	fmt.Println("\n")

	// Access using loop
	for i := 0; i < len(refridgerators); i++ {
		fmt.Println(refridgerators[i])
	}

	fmt.Println("\n")

	// The above code might be error prone, with the condition, we need to always ensure that its less than the length of the array. The best war is as follows:
	for index, refridgerator := range refridgerators {
		fmt.Println(index, refridgerator)
	}

	// Blank identifiers: If we do not want any variable, then we use a blank variable, and its denoted by "_"
	fmt.Println("\n")
	for _, name := range refridgerators {
		fmt.Println(name)
	}

	fmt.Println("\n")
	for index, _ := range refridgerators {
		fmt.Println(index)
	}
}
