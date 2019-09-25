package main

import (
	"fmt"
)

func plusMinus(arr []int32) {
	var positive, negative, zero int32 = 0, 0, 0
	for i := 0; i < len(arr); i ++ {
		if (arr[i] > 0) {
			positive ++
		} else if (arr[i] < 0) {
			negative ++
		} else {
			zero ++
		}
	}
	var percentpositive float32 = float32(positive) / float32(len(arr))
	var percentnegative float32 = float32(negative) / float32(len(arr))
	var percentzero float32 = float32(zero) / float32(len(arr))

	fmt.Printf("%.6f\n", percentpositive)
	fmt.Printf("%.6f\n", percentnegative)
	fmt.Printf("%.6f\n", percentzero)
}

func main() {
	var arr = []int32 {-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}