package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	foo()
	bar()
	end := time.Now()
	fmt.Println("Time taken: ", end.Sub(start))
}

func foo() {
	for i := 0; i < 45; i++ {
		fmt.Println("Foo :", i)
	}
}

func bar() {
	for i := 0; i < 45; i++ {
		fmt.Println("Bar :", i)
	}
}
