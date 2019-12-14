package main

import (
	"fmt"
)

func main() {
	fmt.Println("Executing incrementer")
	c := incrementer()
	fmt.Println("c is: ", c)
	fmt.Println("Executing calculateSum")
	sum := calculateSum(c)
	fmt.Println("sum is: ", sum)
	for value := range sum {
		fmt.Println("Receiving output sum")
		fmt.Println(value)
	}
}

func incrementer() chan int {
	fmt.Println("started incrementer")
	out := make(chan int)
	fmt.Println("channel created in incrementer")
	go func() {
		fmt.Println("Started subprocess 1 for incrementer")
		for i := 0; i < 10; i ++ {
			fmt.Printf("passing %d to the channel\n", i)
			out <- i
		}
		close(out)
		fmt.Println("subprocess 1 ended for incrementer")
	}()
	return out
}

func calculateSum(c chan int) chan int {
	fmt.Println("started calculateSum")
	out := make(chan int)
	fmt.Println("channel created in calculateSum")
	go func() {
		fmt.Println("Started subprocess 2 for calculateSum")
		var sum int
		for value := range c {
			fmt.Printf("receiving %d from the channel\n", value)
			fmt.Printf("adding %d to %d\n", value, sum)
			sum += value
		}
		out <- sum
		close(out)
		fmt.Println("subprocess 2 ended for calculateSum")
	}()
	return out
}
