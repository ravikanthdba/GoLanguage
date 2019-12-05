package mymath

import (
	"fmt"
	"testing"
)

func TestCenteredAvg(t *testing.T) {
	a := []int{1, 4, 6, 8, 100}
	if CenteredAvg(a) != 6 {
		t.Errorf("Expected: 6, Got: %f", CenteredAvg(a))
	}
}

func ExampleCenteredAvg() {
	b := []int{0, 8, 10, 1000}
	fmt.Println(CenteredAvg(b))
	// Output:
	// 9
}

func BenchmarkCenteredAvg(b *testing.B) {
	c := []int{9000, 4, 10, 8, 6, 12}
	for i := 0; i < b.N; i++ {
		CenteredAvg(c)
	}
}
