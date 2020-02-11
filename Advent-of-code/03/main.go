package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func functionality(originalSlice *[]int, slice [][]int) []int {
	fmt.Println("-----------------------------------------------")
	fmt.Println("Now the slice is: ", originalSlice)
	/*
		Requirement:
			1. parse the slice which is multi-splitted
			2. the first item in the slice denotes, which operation to perform
				a) opcode 1: +
				b) opcode 2: *
				c) opcode 99: stop
				d) any other opcode: something is wrong
			3. second item shows the variable in that location of the slice, if the slice is s, and the number given is 20, then the 19th element in the slice to be considered
			4. third item shows the variable in that location of the slice, if the slice is s, and the number given is 28, then the 27th element in the slice to be considered
			5. Based on the operation found out in step 3, perform the activity on the numbers u got in step 3 and 4.
			6. fourth item shows the location where the output of step 5 needs to be stored. IF the slice is s, and the input is 22, then the output needs to be stored in the 21st location.
			7. Ignore the slice, that has number of elements less than 4.
			8. stop the program when you see 99
	*/

	for i := 0; i < len(slice); i++ {
		fmt.Println("slice is: ", slice[i])
		fmt.Println("The original slice is: ", originalSlice)
		fmt.Println("The multiline splitter is: ", slice)
		if slice[i][0] == 99 {
			fmt.Println("It's 99, halting the program now")
			break
		} else {
			switch {
			case len(slice[i]) < 4:
				continue
			case slice[i][0] == 1:
				fmt.Printf("Adding %d and %d\n", (*originalSlice)[slice[i][1]], (*originalSlice)[slice[i][2]])
				(*originalSlice)[slice[i][3]] = (*originalSlice)[slice[i][1]] + (*originalSlice)[slice[i][2]]
			case slice[i][0] == 2:
				fmt.Printf("Multiplying %d and %d\n", (*originalSlice)[slice[i][1]], (*originalSlice)[slice[i][2]])
				(*originalSlice)[slice[i][3]] = (*originalSlice)[slice[i][1]] * (*originalSlice)[slice[i][2]]
			}
			fmt.Println("writing to location: ", slice[i][3])
			fmt.Printf("Replacing the slice value at %d location to %d: \n", slice[i][3], (*originalSlice)[slice[i][3]])
			fmt.Println(" Slice after chaing: ", *originalSlice)
		}
	}
	return *originalSlice
}

// 1,9,10,3,2,3,11,0,99,30,40,50

func multiLineSplitter(slice []int) [][]int {
	/*
		Algorithm to split the slice as per the logic:
			1. Input the slice from the file input.txt
			2. Create a new outerSlice
			3. Iterate through the slice, where u took elements from the file
			4. Create a new inner slice
			5. Until the length of inner slice is 4, keep adding elements to the inner slice.
			6. Once the inner slice is 4, stop adding elements
			7. Add the inner slice to the outer slice.
			8. Remove all the elements that were added from the main slice.
			9. Keep iterating
	*/

	var innerSlice []int
	var outerSlice [][]int

	for len(slice) != 0 {
		if len(innerSlice) < 4 {
			if slice[0] == 99 {
				if len(innerSlice) == 0 {
					/* 1. If the element 99 is the first in the slice, then immediately append it. */
					innerSlice = append(innerSlice, 99)
				} else {
					/* If the element 99 is somewhere in the middle or at the end, then do the following:
					1. append the innerSlice to the outerSlice
					2. Flush the inner slice
					3. Add the element 99 to the inner slice
					*/
					outerSlice = append(outerSlice, innerSlice)
					innerSlice = []int{}
					innerSlice = append(innerSlice, 99)
				}
				/*
					4. Append the innerSlice to the outerSlice
					5. Flush the slice again
				*/
				outerSlice = append(outerSlice, innerSlice)
				innerSlice = []int{}
			} else {
				innerSlice = append(innerSlice, slice[0])
			}
			slice = slice[1:]
			continue
		} else {
			outerSlice = append(outerSlice, innerSlice)
			innerSlice = []int{}
		}
	}
	if len(innerSlice) != 0 {
		outerSlice = append(outerSlice, innerSlice)
	}

	return outerSlice
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
		for _, value := range strings.Split(data.Text(), ",") {
			int, _ := strconv.Atoi(value)
			sliceIntegers = append(sliceIntegers, int)
		}
	}

	fmt.Println("Original slice: ", sliceIntegers)
	fmt.Println("Length of multi line splitter is: ", len(multiLineSplitter(sliceIntegers)))
	iterationValue := 1
	fmt.Println(multiLineSplitter(sliceIntegers))

	for i := 0; i < iterationValue; i++ {
		multiLine := multiLineSplitter(sliceIntegers)
		sliceIntegers = functionality(&sliceIntegers, multiLine[i:])
		fmt.Println("After iteration: ", i, " slice is now: ", sliceIntegers)
		fmt.Println()
	}

	fmt.Println("finally the slice is: ", sliceIntegers)
}
