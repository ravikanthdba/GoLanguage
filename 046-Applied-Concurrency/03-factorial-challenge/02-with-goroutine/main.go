package main

import (
	"fmt"
)

func main() {
	c1 := passData(6)
	c2 := factorial(c1)

	fmt.Println(<-c2)

}

func passData(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i <= n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(c chan int) chan int {
	out := make(chan int)
	var fact int = 1

	go func() {
		for number := range c {
			fact *= number
		}
		out <- fact
		close(out)
	}()
	return out
}
