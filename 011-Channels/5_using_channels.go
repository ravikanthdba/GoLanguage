package main

import (
	"fmt"
)

func main() {

	c := make(chan int)
	fmt.Printf("The type of the channel c is: %T\n", c)
	fmt.Println("The above channel is a bidirectional channel")

	go foo(c)
	go bar(c)

	fmt.Println("The value of <-c is:", <-c)
}

func foo(c chan<- int) {
	c <- 42
}

func bar(c <-chan int) {
	fmt.Println(<-c)
}
