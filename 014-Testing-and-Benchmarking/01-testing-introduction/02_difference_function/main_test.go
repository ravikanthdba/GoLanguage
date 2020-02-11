package main

import (
	"testing"
)

func Test_difference1(t *testing.T) {
	var variable1 int = 100
	var variable2 int = 200

	if difference(variable1, variable2) != 100 {
		t.Errorf("The difference is %d\n, which is not equal to 100", difference(variable1, variable2))
	}

	t.Logf("The difference is %d\n, which is equal to 100", difference(variable1, variable2))
}

func Test_difference2(t *testing.T) {
	var variable1 int = 187
	var variable2 int = 2

	if difference(variable1, variable2) != 185 {
		t.Errorf("The difference is %d\n, which is not equal to 185", difference(variable1, variable2))
	}

	t.Logf("The difference is %d\n, which is equal to 185", difference(variable1, variable2))
}
