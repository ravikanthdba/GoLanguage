package reverse

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := ReverseString("Hello")
	if s != "olleH" {
		t.Errorf("Expected: olleH received: %s\n", s)
	}
}

func TestReverseInteger(t *testing.T) {
	reversed_integer := ReverseInteger(123)
	if reversed_integer != 321 {
		t.Errorf("Expected: 321 received: %d\n", reversed_integer)
	}
}

func ExampleReverseString() {
	fmt.Println(ReverseString("Hello"))
	// Output:
	// olleH
}

func ExampleReverseInteger() {
	fmt.Println(ReverseInteger(98761))
	// Output:
	// 16789
}

func BenchmarkReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseString("Hello")
	}
}

func BenchmarkReverseInteger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseInteger(9876123)
	}
}
