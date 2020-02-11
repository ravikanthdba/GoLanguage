package main

import (
	"fmt"
	"sync"
)

func main() {
	odd_numbers := make(chan int)
	even_numbers := make(chan int)
	fan_in_factorial := make(chan int)

	go getData(odd_numbers, even_numbers)
	go runFactorial(odd_numbers, even_numbers, fan_in_factorial)

	for value := range fan_in_factorial {
		fmt.Println(value)
	}
}

func getData(odd_numbers, even_numbers chan int) {
	for i := 1; i < 20; i++ {
		if i%2 == 0 {
			odd_numbers <- i
		} else {
			even_numbers <- i
		}
	}

	close(odd_numbers)
	close(even_numbers)
}

func runFactorial(odd_numbers, even_numbers, fan_in_factorial chan int) {
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		for value := range odd_numbers {
			fan_in_factorial <- fact(value)
		}
		wg.Done()
	}()

	go func() {
		for value := range even_numbers {
			fan_in_factorial <- fact(value)
		}
		wg.Done()
	}()

	wg.Wait()
	close(fan_in_factorial)
}

func fact(n int) int {
	var fact int = 1

	for i := 1; i <= n; i++ {
		fact *= i
	}

	return fact
}
