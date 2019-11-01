package main

import (
    "fmt"
    "log"
    "math"
)

func main() {
    square_root, err := sqrt(-10);
    if err != nil {
        log.Fatalln(err);
    }
    fmt.Println("Square root of the number is: ", square_root);
}

func sqrt(value float64) (float64, error) {
    if value < 0 {
        errorMessage := fmt.Errorf("The number passed is %f, and since %f < 0 is %v, We cannot calculate square root", value, value, (value < 0)); 
        return 0, errorMessage; 
    }
    return math.Sqrt(value), nil;
}