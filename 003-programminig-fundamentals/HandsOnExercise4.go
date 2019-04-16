/* 
    Write a program to do the following:

    a) Assign an integer to a variable
    b) Prints the integer in Binary, Octal, Decimal and Hexadecimal format
    c) Shifts the binary to the left by 1 digit and assign it to a variable
    d) Print the variable in binary, decimal, octal and Hexa Decimal format
*/



package main

import (
	"fmt"
)

func main() {
	var integer int = 42;
	fmt.Printf("The number %d in binary is %b\n", integer, integer);
	fmt.Printf("The number %d in decimal is %d\n", integer, integer);
	fmt.Printf("The number %d in octal is %o\n", integer, integer);
	fmt.Printf("The number %d in Hexa decimal is %#x\n", integer, integer);
	fmt.Printf("The number %d in unicode is %U\n", integer, integer);

	leftShift := integer << 1
    fmt.Printf("The binary number %b in decimal is %d\n", leftShift, leftShift);
	fmt.Printf("The binary number %b in octal is %o\n", leftShift, leftShift);
	fmt.Printf("The binary number %b in hexadecimal is %#x\n", leftShift, leftShift);
	fmt.Printf("The binary number %b in binary is %b\n", leftShift, leftShift);
	fmt.Printf("The binary number %b in unicode is %U\n", leftShift, leftShift);

}