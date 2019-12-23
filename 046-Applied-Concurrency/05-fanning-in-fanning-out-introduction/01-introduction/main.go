/* Display all the numbers by fanning out odd and even numbers */

package main

import (
	"fmt"
	"sync"
)

func main() {
	odd_channel := make(chan int)
	even_channel := make(chan int)
	fan_in_channel := make(chan int)

	go send(odd_channel, even_channel)
	go receive(odd_channel, even_channel, fan_in_channel)

	for value := range fan_in_channel {
		fmt.Println(value)
	}
}

func send(odd_channel, even_channel chan int) {
	fmt.Println("Launched Subprocess for send")
	for i := 0; i <= 1000; i++ {
		fmt.Println("sending value i to the channel:", i)
		if i%2 != 0 {
			odd_channel <- i
		} else {
			even_channel <- i
		}
	}
	close(odd_channel)
	close(even_channel)
}

func receive(odd_channel, even_channel, fan_int_channel chan int) {
	fmt.Println("Launched Subprocess for receive")
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for value := range odd_channel {
			fmt.Println("receiving odd value from the channel ", value)
			fan_int_channel <- value
		}
		wg.Done()
	}()

	go func() {
		for value := range even_channel {
			fmt.Println("receiving even value from the channel")
			fan_int_channel <- value
		}
		wg.Done()
	}()

	wg.Wait()
	close(fan_int_channel)
}
