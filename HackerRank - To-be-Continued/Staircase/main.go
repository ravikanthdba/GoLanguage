package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 6
	for i := 0; i <= n; i++ {
		fmt.Print(strings.Repeat(" ", (n - i)))
		for j := 0; j < n; j++ {
			if (i + j) >= n {
				fmt.Print("#")
			}
		}
		fmt.Println("")
	}
}
