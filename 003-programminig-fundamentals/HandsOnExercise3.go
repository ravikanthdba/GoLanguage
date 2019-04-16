/* Print TYPED and UNTYPED constants */


/* 

	In Constants, we should not declare the variable as "var <variable_name> <datatype>".
	We should define as : <variable_name> <data_type> = <value>
*/

package main

import (
	"fmt"
)

const (
	a = 43; // Untyped constant
	b int = 65; // Typed Constant

	x = 43.256; // Untyped Constant
	y float32 = 2.13145; // Typed Constant
)

func main() {
	fmt.Printf("%d\n", a);
	fmt.Printf("%d\n", b);

	fmt.Printf("%T\n", a);
	fmt.Printf("%T\n", b);

	fmt.Printf("%f\n", x);
	fmt.Printf("%f\n", y);

	fmt.Printf("%T\n", x);
	fmt.Printf("%T\n", y);
}