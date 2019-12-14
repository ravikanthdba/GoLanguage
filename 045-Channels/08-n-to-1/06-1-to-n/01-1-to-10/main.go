package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	done := make(chan bool)
	n := 10

	fmt.Println("Channel c created with ", c)
	fmt.Println("Channel done created with ", done)


	go func () {
		for i := 0; i < 10000; i ++ {
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i ++ {
		fmt.Println("Launching the subprocess: ", i)
		go func() {
			var process []int
			for value := range c {
				process = append(process, value)
			}
			fmt.Println("process list is now: ", process)
			done <- true
		}()
	}
}
