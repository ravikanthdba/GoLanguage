package main

import "fmt"

func main() {
	s := "i'm sorry dave i can't do that"
	fmt.Println("The original string is: ", s)

	// convert to slice of bytes
	fmt.Println("The string after conversion of slice of bytes: ", []byte(s))

	// convert byte of slice to string
	fmt.Println("After conversion of byte of slice to string is: ", string([]byte(s)))

	// string slice
	fmt.Println("Slicing example 1: ", s[10:])

	// string slice
	fmt.Println("Slicing example 2: ", s[10:22])

	// string slice
	fmt.Println("Slicing example 3: ", s[17:])

	// print rune on each line
	for _, value := range s {
		fmt.Printf("%v - %T - %s \n", rune(value), rune(value), string(rune(value)))
	}
}
