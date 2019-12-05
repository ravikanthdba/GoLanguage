/* Print out the years you are alive */

package main

import (
	"fmt"
)

func main() {
	var birthyear int
	fmt.Println("Enter the Year of Birth: ")
	fmt.Scanf("%d", &birthyear)

	fmt.Printf("%d is the year of birth\n", birthyear)
	fmt.Printf("Your age right now is : %d years\n", (2019 - birthyear))

	fmt.Println("The following are the years you are alive:")
	for birthyear <= 2019 {
		fmt.Println(birthyear)
		birthyear++
	}

}
