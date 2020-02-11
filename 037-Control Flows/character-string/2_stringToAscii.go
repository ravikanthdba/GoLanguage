package main

import (
	"fmt"
)

func main() {
	var word string = "hello"
	for letter := 0; letter < len(word); letter++ {
		fmt.Println(letter, "------", word[letter])
	}
}
