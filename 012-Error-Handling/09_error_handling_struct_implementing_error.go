package main

import (
    "fmt"
    "log"
)

type mathError struct {
    latitude string
    longitude string
    err error
}

func (m mathError) Error() string {
    return fmt.Sprintf("Error occurred with %v, %v and %v", m.latitude, m.longitude, m.err);
}

func main() {
    _, err := sqrt(-10.22);
    if err != nil {
        log.Println(err);
    }
}

func sqrt(value float64) (float64, error) {
    if value < 0 {
        errorMessage := fmt.Errorf("value %f is less than 0", value);
        return 0, mathError{"50.1234N", "99.6758W", errorMessage};
    }
    return 42, nil;
}