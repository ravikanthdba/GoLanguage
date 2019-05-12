/*

Pointers are used, when we want to pass a huge chunk of data movement, and if you want to optimize performance.
In the below example, we see when we pass a variable inito a function, and try changee that variable

*/

package main

import (
	"fmt"
)

func main() {

	a := 0;
	b := 1;

	fmt.Println("The value of a from main function is ", a);
	fmt.Println("The value of b from main function is ", b);

	fmt.Println("The address of a from main function is ", &a);
	fmt.Println("The address of b from main function is ", &b);

	foo(a, &b);

	fmt.Println("The value of a from main function is ", a);
	fmt.Println("The value of b from main function is ", b);

	fmt.Println("The address of a from main function is ", &a);
	fmt.Println("The address of b from main function is ", &b);	

}

func foo(x int, y *int) {

	fmt.Println("The value of x from foo before change", x);
	x = 42;
	fmt.Println("The value of x from foo after  change", x);


	fmt.Println("The value of y from foo before change", y);
	*y = 99;
	fmt.Println("The value of y from foo after  change", y);

}