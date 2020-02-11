/* Sort the slice of Integer or Slice of Strings */

package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println("Unsorted: ", xi)

	sort.Ints(xi)
	fmt.Println("Sorted:", xi)

	xs := []string{"Ravikanth", "Viraj", "NAgabhushanam", "Swarna Latha", "Bhargavi", "Divya", "Kishore"}
	fmt.Println("Unsorted:", xs)
	sort.Strings(xs)
	fmt.Println("Sorted: ", xs)
}
