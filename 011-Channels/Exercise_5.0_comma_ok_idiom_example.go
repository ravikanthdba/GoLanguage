// show comma ok idiom starting with https://play.golang.org/p/YHOMV9NYKK

package main

import (
	"fmt"
)

func main() {
	channel := make(chan int)

	go func() {
		channel <- 42
	}()

	value, ok := <-channel
	fmt.Println("value: ", value, " and status ok is: ", ok)

	close(channel)

	value, ok = <-channel
	fmt.Println("value: ", value, " and status ok is: ", ok)
}
