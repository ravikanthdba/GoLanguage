/*
Get this code running:
https://play.golang.org/p/oB-p3KMiH6
*/

package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go send(c)
	go receive(c)

	fmt.Println("The data for c is now: ", <-c)
}

func send(c chan<- int) {
	c <- 42
}

func receive(c <-chan int) {
	fmt.Println(<-c)
}
