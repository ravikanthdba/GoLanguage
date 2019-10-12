// Unlike arrays, slice can grow in size randomly.

/*
	array declaration: [3]int {0, 1, 2}
    slice declaration: []int {0, 1, 2}
*/

package main

import (
	"fmt"
)

func main() {

	var wholenumbers []int
	wholenumbers = []int{1,2,3,4,5}
	fmt.Println(wholenumbers)
	fmt.Println("\n")

		// Declare a variable that holds a slice of ints.
		var wholeNumbers []int
		// Create a new slice of 5 elements and assign it to the variable.
		wholeNumbers = []int{1,2,3,4,5}
		// Assign 42 to the second element.
		wholeNumbers[2] = 42
		fmt.Printf("%#v\n", wholeNumbers) // => []int{0, 42, 0, 0, 0}
	

	variables := []string {"Ravikanth", "Garimella", "Bhargavi", "Bharatula", "Viraj", "Nandan"}
	fmt.Println(variables)

	if (variables == nil) {
		fmt.Println("variables is null")
	} else {
		fmt.Println("Variables is not null")
	}

	var mySlice []int
	if (mySlice == nil) {
		fmt.Println("mySlice is null")
	}
	fmt.Println("The slice is now: ", mySlice)
	fmt.Printf("%#v\n", mySlice)	


	// Make a new slice with the built-in make function
	var mySlicemakeint []int
	mySlicemakeint = make([]int, 3)
	if (mySlicemakeint == nil) {
		fmt.Println("mySlicemakeint is null")
	} else {
		fmt.Println("mySlicemakeint is not null")
	}
	fmt.Println("The slice is now", mySlicemakeint)
	fmt.Printf("%#v\n", mySlicemakeint)

	var mySlicemakechar []string
	mySlicemakechar = make([]string, 5)
	if (mySlicemakechar == nil) {
		fmt.Println("mySlicemakechar is null")
	} else {
		fmt.Println("mySlicemakechar is not null")
	}
	fmt.Println(mySlicemakechar)
	fmt.Printf("%#v\n", mySlicemakechar)


	// Primes array using slices
	primeNumbers := make([]int, 4)
	if (primeNumbers == nil) {
		fmt.Println("primeNumbers is null")
	} else {
		fmt.Println("primeNumbers is not null")
	}
	fmt.Println(primeNumbers)
	fmt.Printf("%#v\n", primeNumbers)	
	primeNumbers[0] = 1
	primeNumbers[3] = 4
	fmt.Println(primeNumbers)
	fmt.Printf("%#v\n", primeNumbers)	

	fmt.Println("\n")
	// Add more variables using the append function.
	fmt.Printf("%#v\n", primeNumbers)
	primeNumbers = append(primeNumbers, 5)	
	fmt.Printf("%#v\n", primeNumbers)

	fmt.Println("\n")
	nameSlices := []string{"Nagabhushanam", "Swarna", "Divya", "Kishore", "Atharv","Ravikanth", "Bhargavi", "Viraj"}
	for name := 0; name < len(nameSlices); name++ {
		fmt.Println(name, nameSlices[name])
	}
	fmt.Println("\n")
	for index, names := range nameSlices {
		fmt.Println(index, names)
	}

	// Exercise - https://play.golang.org/p/wfeezKx0S2v
	fmt.Println("\n")

	myArray := [5]int{0,1,2,3,4}
	fmt.Printf("%#v\n", myArray[2])
	fmt.Printf("%#v\n", myArray[2:]) // Generates a slice of integers
}