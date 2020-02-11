package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan int, 10)
	go send(c)
	go receive(c)
	time.Sleep(time.Second)
	end := time.Now()
	fmt.Println("Time Taken:", end.Sub(start))
}

func send(s chan int) {
	for i := 0; i < 100; i++ {
		fmt.Println("The value sent is: ", i)
		fmt.Printf("Sending %v to the channel\n", s)
		s <- i
	}
}

func receive(r chan int) {
	for {
		fmt.Println("The value received is: ", <-r)
		fmt.Printf("Receiving %v from the channel\n", r)
	}
}
