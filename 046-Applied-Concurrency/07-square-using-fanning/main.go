package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3, 10)

	c1 := sq(in)
	c2 := sq(in)
	c3 := sq(in)

	for data := range merge(c1, c2, c3) {
		fmt.Println(data)
	}

}

func gen(n ...int) chan int {
	out := make(chan int)
	go func() {
		for value := range n {
			out <- n[value]
		}
		close(out)
	}()
	return out
}

func sq(c chan int) chan int {
	out := make(chan int)
	go func() {
		for value := range c {
			out <- value * value
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
