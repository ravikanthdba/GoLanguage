package main

import (
	"fmt"
)

func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
	var applecount int = 0
	var orangecount int = 0
	for apple := 0; apple < len(apples); apple++ {
		if a+apples[apple] >= s && a+apples[apple] <= t {
			applecount += 1
		}
	}

	for orange := 0; orange < len(oranges); orange++ {
		if b+oranges[orange] >= s && b+oranges[orange] <= t {
			orangecount += 1
		}
	}

	fmt.Println(applecount, orangecount)
}

func main() {
	// var s, t, a, b int32
	var apples = []int32{-2, 2, 1}
	var oranges = []int32{5, -6}
	countApplesAndOranges(7, 11, 5, 15, apples, oranges)
}
