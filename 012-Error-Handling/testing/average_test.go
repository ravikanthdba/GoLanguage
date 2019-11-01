/*
We are going to learn about testing next. For this exercise, take a moment and see how much you can figure out about testing by reading the testing documentation & also Caleb Doxsey’s article on testing. See if you can get a basic example of testing working. 
*/

package main

import (
    "testing"
)

func TestAverage(t *testing.T) {
    var v float64;
    v = average([]float64{1, 2, 3, 4, 5, 6});
    if v != 3.5 { 
        t.Error("Expected 3.5, got ", v);
    }
}
