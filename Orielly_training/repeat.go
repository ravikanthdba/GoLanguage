package main

import (
	"fmt"
)

func repeat(word string, times int) {
	for i := 0; i < times; i ++ {
		fmt.Println(word);
	}
}

func main() {
	repeat("Hello", 4);
}