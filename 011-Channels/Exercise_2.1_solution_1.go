package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go send(c)
	go receive(c)

	fmt.Println("Finishing the program..")
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
