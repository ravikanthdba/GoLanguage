package main

import (
	"fmt"
)

func main() {
	send := sendData(4, 5, 6)
	receive := sqr(send)

	fmt.Println(<- receive)
	fmt.Println(<- receive)
	fmt.Println(<- receive)
}

func sendData(n ...int) <- chan int {
	out := make(chan int)
	go func() {
		for value := range n {
			out <- n[value]
		}
		close(out)
	}()
	return out
}

func sqr(c <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		for value := range c {
			out <- value * value
		}
		close(out)
	}()
	return out
}
