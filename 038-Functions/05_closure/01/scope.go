package main

import (
	"fmt"
) // accessible throughout the file

func main() {
	x := 42                             // scope of x is throughout the main function
	fmt.Println("From main, x is: ", x) // x is accessible here
	{
		fmt.Println("From inside subfunction, x is: ", x) // x is accessible here
		y := 50                                           // y is limited only to this block
		fmt.Println("From inside subfunction, y is: ", y) // y is accessible here
	}
	//  fmt.Println("From main, y is: ", y); // y is not accessible here, as y is outside of the inner block
}
