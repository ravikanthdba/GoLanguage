package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go send(c)
	go receive(c)
	time.Sleep(time.Second)
}

func send(s chan int) {
	for i := 0; i < 20; i ++ {
		fmt.Println("The value sent is: ", i)
		fmt.Printf("Sending %v to the channel\n", s)
		s <- i
	}
}

func receive(r chan int) {
	for {
		fmt.Println("The value received is: ", <- r)
		fmt.Printf("Receiving %v from the channel\n", r)
	}
}