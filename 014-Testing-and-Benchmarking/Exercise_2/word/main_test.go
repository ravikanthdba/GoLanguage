package word

import (
	"fmt"
	"testing"
)

func TestCount(t *testing.T) {
	sentence := "Here is the new sentence"
	countWords := Count(sentence)
	if countWords != 5 {
		t.Errorf("Expected: 5, returned: %d", countWords)
	}
}

func ExampleCount() {
	sentence := "Here is the new sentence"
	fmt.Println(Count(sentence))
	// Output:
	// 5
}

func BenchmarkCount(t *testing.B) {
	sentence := "Here is the new sentence"
	for i := 0; i < t.N; i++ {
		Count(sentence)
	}
}
