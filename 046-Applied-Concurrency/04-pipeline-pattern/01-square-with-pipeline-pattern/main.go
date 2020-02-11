package main

import (
	"fmt"
)

func main() {
	c := gen(2, 3)
	sq := sqr(c)
	fmt.Println(<-sq)
	fmt.Println(<-sq)

}

func gen(c ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, number := range c {
			out <- number
		}
		close(out)
	}()
	return out
}

func sqr(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	return out
}
