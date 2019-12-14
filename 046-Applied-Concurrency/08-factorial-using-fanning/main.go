package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen()

	c0 := fact(in)
	c1 := fact(in)
	//c2 := fact(in)
	//c3 := fact(in)
	//c4 := fact(in)
	//c5 := fact(in)
	//c6 := fact(in)
	//c7 := fact(in)
	//c8 := fact(in)
	//c9 := fact(in)

	for value := range merge(c0, c1) {
		fmt.Println(value)
	}

}

func gen() chan int {
	out := make(chan int)

	go func() {
		for i := 0; i < 10; i ++ {
			for j := 1; j <= 2; j ++ {
				out <- j
			}
		}
		close(out)
	}()

	return out
}

func fact(c chan int) chan int {
	var fact int = 1
	out := make(chan int)

	go func() {
		for value := range c {
			for i := 1; i <= value; i ++ {
				fact *= i
			}
			out <- fact
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

		go func() {
			wg.Wait()
			close(out)
		}()
	}

	return out
}
