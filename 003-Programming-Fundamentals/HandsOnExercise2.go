/* Using the below operators, write expressions and assign values to variables:
   ==
   <=
   >=
   !=
   >
   <
*/

package main

import "fmt"

func main() {
	x := (42 == 43)
	y := (42 >= 43)
	z := (42 <= 43)
	w := (42 != 43)
	r := (42 < 43)
	s := (42 > 43)

	fmt.Println("The value of x is ", x)
	fmt.Println("The value of y is ", y)
	fmt.Println("The value of z is ", z)
	fmt.Println("The value of w is ", w)
	fmt.Println("The value of r is ", r)
	fmt.Println("The value of s is ", s)

}
