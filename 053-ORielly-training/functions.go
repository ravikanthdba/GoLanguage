package main

import (
	"fmt"
)

func printHi() {
	fmt.Println("Hello")
}

func printEvent(number int) {
	if number%2 == 0 {
		fmt.Printf("%d is even\n", number)
	} else {
		fmt.Printf("%d is odd\n", number)
	}
}

func sayNtimes(word string, t int) {
	for i := 0; i < t; i++ {
		fmt.Printf("%s\n", word)
	}
}

func main() {
	printHi()

	number := 11
	printEvent(number)

	sayNtimes("how are you", 5)
}
