package main

import (
	"fmt"
	"sort"
)

const (
	max = 4
)

func miniMaxSum(arr []int) (int, int) {
	var outputSum []int

	for i := 0; i <= max; i++ {
		var sum int = 0
		for j := 0; j <= max; j++ {
			if i != j {
				sum += arr[j]
			}
		}
		outputSum = append(outputSum, sum)
	}

	sort.Ints(outputSum)

	var min int = outputSum[0]
	var max int = outputSum[len(outputSum)-1]

	return min, max

}

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(miniMaxSum(arr))
}
