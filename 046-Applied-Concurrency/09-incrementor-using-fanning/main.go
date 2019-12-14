package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	c0 := incrementor(in)
	c1 := incrementor(in)
	c2 := incrementor(in)
	c3 := incrementor(in)
	c4 := incrementor(in)
	c5 := incrementor(in)
	c6 := incrementor(in)
	c7 := incrementor(in)
	c8 := incrementor(in)
	c9 := incrementor(in)

	var sum int

	for data := range merger(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		sum += data
	}

	fmt.Println("The sum is: ", sum)
}

func gen() chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10000; i ++ {
			for j := 0; j < 1000; j++ {
				out <- 1
			}
		}
		close(out)
	}()
	return out
}

func incrementor(c chan int) chan int {
	out := make(chan int)
	var sum int = 0

	go func() {
		for value := range c {
			sum += value
		}
		out <- sum
		close(out)
	}()

	return out
}

func merger(n ...chan int) chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	wg.Add(len(n))

	for _, value := range n {
		go func(ch chan int) {
			for value := range ch {
				out <- value
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
