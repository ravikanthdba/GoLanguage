package main

import (
	"fmt"
)

func main() {
	for data := range sq(sendData(2, 3, 4, 5)) {
		fmt.Println(data)
	}

}

func sendData(n ...int) chan int {
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