/*

Hands-on exercise #6
write a program that puts 100 numbers to a channel and pull the numbers off the channel and print them

*/

package main

import (
	"fmt"
)

func main() {
	channel := send()
	fmt.Println(channel)
	fmt.Printf("%T\n", channel)
	receive(channel)
	fmt.Println("Exiting the program..")
}

func send() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func receive(channel <-chan int) {
	for value := range channel {
		fmt.Println("Value is :\t", value)
		fmt.Println("Channel is :\t", channel)
		fmt.Println("<- Channel is :\t", <-channel)
	}
}
