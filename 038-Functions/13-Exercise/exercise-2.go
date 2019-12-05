package main

import (
	"fmt"
)

func main() {
	h := func(n int) (int, bool) {
		return n / 2, n%2 == 0
	}

	value, checkBoolean := h(10)
	fmt.Println(value, checkBoolean)
}
