package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	c0 := fact(in)
	c1 := fact(in)
	c2 := fact(in)
	c3 := fact(in)
	c4 := fact(in)
	c5 := fact(in)
	c6 := fact(in)
	c7 := fact(in)
	c8 := fact(in)
	c9 := fact(in)

	var y int
	for value := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		y++
		fmt.Println(y, "\t", value)
	}

}

func gen() chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 100000; i++ {
			for j := 1; j <= 10; j++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func fact(c chan int) chan int {
	out := make(chan int)

	go func() {
		for data := range c {
			out <- factorial(data)
		}
		close(out)
	}()

	return out
}

func merge(n ...chan int) chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(n))

	for _, value := range n {
		go func(ch chan int) {
			for c := range ch {
				out <- c
			}
			wg.Done()
		}(value)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func factorial(n int) int {
	var fact int = 1

	for i := 1; i <= n; i++ {
		fact *= i
	}

	return fact
}
