/*
fanout - in the below code, the numbers are split into 2 halves, odd_numbers and even_numbers.
fanin - then the two lists are taken and merged, its called fan in.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	odd_channel := make(chan int)
	even_channel := make(chan int)
	fanin_channel := make(chan int)

	go send(odd_channel, even_channel)
	go receive(odd_channel, even_channel, fanin_channel)

	for value := range fanin_channel {
		fmt.Println(value)
	}

	fmt.Println("Exiting the program..")
}

func send(odd_channel, even_channel chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even_channel <- i
		} else {
			odd_channel <- i
		}
	}
	close(odd_channel)
	close(even_channel)
}

func receive(odd_channel, even_channel <-chan int, fanin_channel chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for v := range even_channel {
			fanin_channel <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd_channel {
			fanin_channel <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin_channel)
}
