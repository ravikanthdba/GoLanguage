package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(boring("xyz"), boring("yzx")) // 1. This prints the address of channels of xyz and yzx. o/p is 0xc0000a6000 0xc0000a6060, and this gets passed to fanIn function.
	c := fanIn(boring("xyz"), boring("yzx"))  // 2. fanIn returns a channel. Into the fanIn function, we pass two parameters, which are functions boring("xyz") and boring("yzx")
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("Exiting")
}

func boring(msg string) <-chan string { // 3. boring takes in the message xyz and yzx as input.
	c := make(chan string) // 4. a channel is created whose name is "c"
	go func() {            // 5. this is being put in a channel.
		for i := 0; ; i++ { // 6. infinite loop
			c <- fmt.Sprintf("%s %d", msg, i)                            // 7. prints message and the counter "i"
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond) // 8. Sleep for a few seconds.
		}
	}()
	return c // Returns the channel of c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)

	go func() {
		for {
			fmt.Println("From the fanin function 1st goroutine input1 = :", input1, "<- input 1 is ", <-input1)
			c <- <-input1
		}
	}()

	go func() {
		for {
			fmt.Println("From the fanin function 1st goroutine input2 = :", input2, "<- input 2 is ", <-input2)
			c <- <-input2
		}
	}()

	fmt.Println("c is now: ", c, " and <- c = ", <-c)
	return c
}
