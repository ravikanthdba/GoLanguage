/*
	chan T - is a bi-directional channel with type T
	chan <- T - only can send data of type T to the channel - send-only channel
	<- chan T - only can receive data of type T from the channel - receive-only channel
*/

package main

import (
	"fmt"
)

func main() {
	c := incrementer()
	cSum := puller(c)
	for value := range cSum {
		fmt.Println(value)
	}
}

func incrementer() <-chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
