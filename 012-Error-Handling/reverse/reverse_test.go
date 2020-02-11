package myreversefunction

import (
	"testing"
)

func TestReverse(t *testing.T) {
	var value string = "Hello"
	reversed_string := Reverse(value)

	if reversed_string != "olleH" {
		t.Error("Expected olleH and got ", reversed_string)
	}
}
