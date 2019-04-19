/* 

Using a COMPOSITE LITERAL: 
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position. 
Range over the array and print the values out. 
Using format printing
print out the TYPE of the array

*/

package main

import (
	"fmt"
)

func main() {
	array := [5]int{20,30,40,50,60};
	fmt.Println("The elements of the array are: ", array);

	for value := range array {
		fmt.Println(array[value]);
	}


	fmt.Printf("The type of Array is %T\n", array);

}