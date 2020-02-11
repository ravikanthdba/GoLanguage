package main

import (
	"fmt"
)

func main() {
	for value := range factorial(sendData(5)) {
		fmt.Println(value)
	}
}

func sendData(n int) <-chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(c <-chan int) <-chan int {
	out := make(chan int)

	var factorial int = 1

	go func() {
		for value := range c {
			factorial *= value
		}
		out <- factorial
		close(out)
	}()
	return out
}
