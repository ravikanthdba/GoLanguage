/*
   write a program that launches 10 goroutines each goroutine adds 10 numbers to a channel pull the numbers off the channel and print them
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	channel := make(chan int)
	const maxGoroutines = 10

	for goroutine := 0; goroutine < maxGoroutines; goroutine++ {
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Println("Adding ", strconv.Itoa(i), " to the channel.")
				channel <- i
			}
		}()
	}

	for value := 0; value < 100; value++ {
		fmt.Println(<-channel)
	}

	fmt.Println("Exiting..!")
}
