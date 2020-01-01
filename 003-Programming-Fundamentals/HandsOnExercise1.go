/* Go Program to print a number in decimal, binary, octal and Hexadecimal */

package main

import (
	"fmt"
)

func main() {
	var integer int
	fmt.Println("Enter the number in integer")
	fmt.Scanf("%d", &integer)

	fmt.Printf("The number %d in decimal                  is : %d\n", integer, integer)
	fmt.Printf("The number %d in binary                   is : %b\n", integer, integer)
	fmt.Printf("The number %d in octal                    is : %o\n", integer, integer)
	fmt.Printf("The number %d in Hexa Decimial Lower Case is : %#x\n", integer, integer)
	fmt.Printf("The number %d in Hexa Decimial Upper Case is : %#X\n", integer, integer)
	fmt.Printf("The number %d in Unicode Format           is : %U\n", integer, integer)
}
