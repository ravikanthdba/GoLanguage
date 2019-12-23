package main

import (
	"fmt"
)

func main() {
	in := gen()
	f := factorial(in)

	for result := range f {
		fmt.Println(result)
	}
}

func gen() chan int64 {
	out := make(chan int64)
	var j int64

	go func() {
		for i := 0; i < 10; i++ {
			for j = 6; j < 160; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func factorial(c chan int64) chan int64 {
	out := make(chan int64)

	go func() {
		for value := range c {
			out <- fact(value)
		}
		close(out)
	}()

	return out
}

func fact(n int64) int64 {
	var fact int64 = 1
	for i := n; i > 0; i-- {
		fact *= i
	}
	return fact
}
