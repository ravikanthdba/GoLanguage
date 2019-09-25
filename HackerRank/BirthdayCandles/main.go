package main

import (
	"fmt"
	"sort"
)

func main() {
	var arr = []int {4, 4, 1, 2}
	var count = 0
	sort.Ints(arr)
	var maxValue int = arr[len(arr) - 1]
	for i := 0; i <= len(arr) - 1; i++ {
		if arr[i] == maxValue {
			count += 1
		}
	}
	fmt.Println(count)
}