/*
   Rune is a single character.
   rune is represented by ''
   string is represented by "" or ``
   string is different than runes
   alias for int32

*/

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%T\n", 'a')         // Prints the data type of rune, and its int32.
	fmt.Println('a')                // Prints 97, as the ASCII code of 'a' is 97.
	fmt.Println(string('a'))        // 'a' is a rune, and it gets converted to string using the string function. Hence this prints string
	fmt.Printf("%T\n", string('a')) // This prints string as it is converted to string using the string function
	fmt.Println(string(97))         // string value of 97 is 'a', as per the ASCII values
	fmt.Println(int('a'))           //Prints 97 as the ascii value of 'a' is 97
}
