package main

import (
	"fmt"
)

func main() {
	var i int

	for {
		fmt.Println(i)
		i++
		if i == 99 {
			break
		} else {
			continue
		}
	}
}
