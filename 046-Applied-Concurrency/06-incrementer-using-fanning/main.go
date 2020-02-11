package main

import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	fan_in_channel := make(chan int)

	go send(chan1, chan2)
	go receive(chan1, chan2, fan_in_channel)

	for value := range fan_in_channel {
		fmt.Println(value)
	}
}

func send(chan1, chan2 chan int) {
	for i := 0; i < 2000; i++ {
		if i%2 == 0 {
			chan1 <- 1
		} else {
			chan2 <- 1
		}
	}

	close(chan1)
	close(chan2)
}

func receive(chan1, chan2, fan_in_channel chan int) {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for value := range chan1 {
			fan_in_channel <- sumFunction("Foo", value)
		}
		wg.Done()
	}()

	go func() {
		for value := range chan2 {
			fan_in_channel <- sumFunction("Bar", value)
		}
		wg.Done()
	}()

	wg.Wait()
	close(fan_in_channel)
}

var sum int

func sumFunction(s string, n int) int {
	fmt.Println("Adding 1 through the function: ", s)
	sum += n
	return sum
}
