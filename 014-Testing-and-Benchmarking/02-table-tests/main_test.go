package main

import (
	"fmt"
	"testing"
)

func Test_summation(t *testing.T) {
	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{21, 21}, 42},
		test{[]int{2, 21}, 23},
		test{[]int{-1, 0, 1}, 0},
		test{[]int{21, 30}, 51},
	}

	fmt.Printf("%+v\n", tests)

	for _, variable := range tests {
		sum := summation(variable.data...)
		if sum != variable.answer {
			t.Error("Expected answer is: ", variable.answer, " but returned answer is: ", sum)
		}
	}
}
