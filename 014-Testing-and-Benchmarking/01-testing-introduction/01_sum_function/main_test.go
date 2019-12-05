package main

import (
	"testing"
)

func Test_summingNumbers(t *testing.T) {
	sum := summingNumbers(6, 7)
	if sum != 13 {
		t.Errorf("Expected: 13, and returned: %d\n", sum)
	}
	t.Logf("Expected: 13 and returned: %d\n, and hence passed the test..", sum)
}

func Test_summingNumbers_new(t *testing.T) {
	sum := summingNumbers(6, 7, 10, 13, 14)
	if sum != 50 {
		t.Errorf("Expected: 50, and returned: %d\n", sum)
	}
	t.Logf("Expected: 50 and returned: %d\n, and hence passed the test..", sum)
}
