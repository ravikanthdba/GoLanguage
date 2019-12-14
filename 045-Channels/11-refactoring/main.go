package main

import (
	"fmt"
)

func main() {
	c := incrementer()
	for n := range puller(c) {
		fmt.Println(n)
	}
}

func incrementer() <- chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i ++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		var sum int
		for value := range c {
			sum += value
		}
		out <- sum
		close(out)
	}()
	return out
}
