/*
Starting with this code, sort the []int and []string for each person.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println("Unsorted Integers:", xi);
	sort.Ints(xi)
	fmt.Println("Sorted Integers", xi);
	fmt.Println("\n")
	fmt.Println("-------------------------------")
	fmt.Println("\n")
	fmt.Println("Unsorted Strings:", xs);
	sort.Strings(xs)
	fmt.Println("Sorted Strings", xs);


}