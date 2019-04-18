/* Print out the years you are alive with only for */

package main

import (
	"fmt"
)


func main() {

	var birthyear int;
	fmt.Println("Enter the year of birth: ");
	fmt.Scanf("%d", &birthyear);


	fmt.Println("The following are the years you are alive:")
	for {

		if birthyear > 2019 {
			break;
		}


		fmt.Println(birthyear);
		birthyear ++;
	}
}