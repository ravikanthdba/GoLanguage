package main

import (
	"fmt"
)


func breakingRecords(scores []int) []int {
	var min, max int
	var minCounter, maxCounter int = 0, 0
	var output []int

	for i := 0; i < len(scores); i ++ {
		if i == 0 {
			min = scores[i]
			max = scores[i]
		} else {
			if (scores[i] < min) {
				min = scores[i]
				minCounter ++
			} else if (scores[i] > max) {
				max = scores[i]
				maxCounter ++
			} 
		}
	}
	output = append(output, minCounter)
	output = append(output, maxCounter)
	return output
}

func main() {
	// var x = []int {3, 4, 21, 36, 10, 28, 35, 5, 24, 42}
	// var x = []int {12, 24, 10, 24}
	var x = []int {10, 5, 20, 20, 4, 5, 2, 25, 1}
	fmt.Println(breakingRecords(x))
}