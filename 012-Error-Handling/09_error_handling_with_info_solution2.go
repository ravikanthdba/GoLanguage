package main

import (
    "errors"
    "math"
    "log"
    "fmt"
)

var errorMessage = errors.New("Square root of a negative number is not possible.")

func main() {
    square_root, err := sqrt(-10);
    if err != nil {
        fmt.Printf("The format of the variable err is %T\n", errorMessage);
        log.Fatalln("The error is now: ", errorMessage);
    }
    fmt.Println("The square root of the number is: ", square_root);
}

func sqrt(number float64) (float64, error) {
    square_root := math.Sqrt(number);
    if number < 0 {
        return 0, errorMessage;
    }
    return square_root, nil;
}