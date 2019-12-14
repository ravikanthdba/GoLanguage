package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	c0 := sumNumber(in)
	c1 := sumNumber(in)
	c2 := sumNumber(in)
	c3 := sumNumber(in)
	c4 := sumNumber(in)
	c5 := sumNumber(in)
	c6 := sumNumber(in)
	c7 := sumNumber(in)
	c8 := sumNumber(in)
	c9 := sumNumber(in)

	var sum int

	for value := range merge(c0, c1, c2, c3, c4, c5, c6, c7, c8, c9) {
		sum += value
	}

	fmt.Println("The sum is: ", sum)
}

func gen() chan int {
	out := make(chan int)

	go func() {
		for i := 1; i <= 100; i++ {
			for j := 1; j <= 10; j ++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func sumNumber(c chan int) chan int {
	out := make(chan int)
	var sum int

	go func() {
		for value := range c {
			sum += value
		}
		out <- sum
		close(out)
	}()

	return out
}

func merge(n ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(n))

		for _, value := range n {
			go func(number chan int) {
				for value := range number {
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
