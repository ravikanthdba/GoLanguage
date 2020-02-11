package main

import (
	"fmt"
)

func main() {
	func() { // Since there is no function name (identifier), it's called "Anonymous function"
		fmt.Println("I am a self executing anonymous function")
	}() // Since it's executing, on its own, by calling in the "()", its called self executing.
}
