package main

import (
	"fmt"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	var n int = 10000
	var counter int = 0

	for counter < n {
		if x1 != x2 {
			x1 += v1
			x2 += v2
		} else {
			return "YES"
		}
		counter += 1
	}
	return "NO"
}

func main() {
	var x1 int32 = 1
	var v1 int32 = 2

	var x2 int32 = 2
	var v2 int32 = 1

	fmt.Println(kangaroo(x1, v1, x2, v2))

}
