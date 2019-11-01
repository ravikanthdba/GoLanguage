package main

import (
    "testing"
)

func TestMaximum(t *testing.T) {
    var v int;
    v = maximum([]int{2,3,4});
    if v != 4 {
        t.Error("Wrong maximum value, please try again");
    }
}