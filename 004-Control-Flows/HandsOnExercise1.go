/* Print every number from 1 to 10000*/

package main

import (
	"fmt"
)


func main() {

	fmt.Println("Option 1:")

	for i := 0; i <= 10000; i ++ {
		fmt.Println(i);
	}


	fmt.Println("Option 2:")

	var x int32;
	for {
		if x > 10000 {
			break;
		}

		fmt.Println(x);
		x ++;
	}

}



