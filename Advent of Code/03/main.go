package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
	Algorithm to split the slice as per the logic:
		1. Input the slice from the file input.txt
		2. Create a new outerSlice
		3. Iterate through the slice, where u took elements from the file
		4. Create a new inner slice
		5. Until the length of inner slice is 4, keep adding elements to the inner slice.
		5. Once the inner slice is 4, stop adding elements
		6. Add the inner slice to the outer slice.
		7. Remove all the elements that were added from the main slice.
		8. Keep iterating
 */

func splitter(s []int) {
	var outerSlice [][]int

	for i := 0; i < len(s) - 1; i ++ {
		var innerSlice []int
		if s[i] == 99 {
			innerSlice = append(innerSlice, 99)
			s = s[1:]
		} else {
			for j := 0; j < 4; j ++ {
				innerSlice = append(innerSlice, s[j])
			}
		}

		outerSlice = append(outerSlice, innerSlice)
		s = s[4:]
	}

	fmt.Println(outerSlice)
}

func main() {
	var sliceIntegers []int

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	data := bufio.NewScanner(file)
	for data.Scan() {
		for _, value := range strings.Split(data.Text(),",") {
			int, _ := strconv.Atoi(value)
			sliceIntegers = append(sliceIntegers, int)
		}
	}

	splitter(sliceIntegers)
}
