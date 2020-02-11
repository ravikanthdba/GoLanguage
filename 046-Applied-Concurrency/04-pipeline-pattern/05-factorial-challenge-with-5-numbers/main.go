package main

import (
	"fmt"
)

func main() {
	c := sendData(2, 3, 4, 5, 6, 10)
	factorial := fact(c)
	for result := range factorial {
		fmt.Println(result)
	}
}

func sendData(n ...int) chan int {
	out := make(chan int)

	go func() {
		for number := range n {
			out <- n[number]
		}
		close(out)
	}()

	return out
}

func fact(c chan int) chan int {
	out := make(chan int)

	go func() {
		for value := range c {
			var fact int = 1
			for i := 1; i <= value; i++ {
				fact *= i
			}
			out <- fact
		}
		close(out)
	}()
	return out
}
