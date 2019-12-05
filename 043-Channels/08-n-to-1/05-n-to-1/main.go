// This is n functions running and writing into 1 channel.

package main

import (
	"fmt"
)

func main() {
	n := 10
	c := make(chan int)
	done := make(chan bool)

	fmt.Printf("Channel c created with %d address\n", c)
	fmt.Printf("Channel done created with %d address\n", done)

	for i := 0; i < n ; i++ {
		fmt.Println("Launching the subprocess:", i)
		go func() {
			for j := 0; j < 10; j ++ {
				fmt.Printf("Adding the value %d to the channel\n", j)
				c <- j
			}
			fmt.Println("Setting done to true for subprocess", i)
			done <- true
		}()
	}

	go func() {
		for i := 0; i < 10; i ++ {
			fmt.Printf("Clearing the subprocess: %d\n", i)
			<- done
		}
		close(c)
	}()

	for value := range c {
		fmt.Println("The value of c is now: ", value)
	}

}
