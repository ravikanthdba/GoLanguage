package dog

import (
    "testing"
    "fmt"
)

func TestYears(t *testing.T) {
    var humanYears int = 20;
    if Years(humanYears) != 140 {
        t.Errorf("Expected: 140, Got: %d", Years(humanYears));
    }
}

func TestYearsTwo(t *testing.T) {
    var number int = 10;
    if YearsTwo(number) != 70 {
        t.Errorf("Expected: 70, Got: %d", YearsTwo(number));
    }
}

func ExampleYears() {
    fmt.Println(Years(20));
    // Output:
    // 140
}

func ExampleYears_1() {
    fmt.Println(Years(30));
    // Output:
    // 210
}

func ExampleYearsTwo() {
    fmt.Println(YearsTwo(10));
    // Output:
    // 70
}

func ExampleYearsTwo_1() {
    fmt.Println(YearsTwo(15));
    // Output:
    // 105
}


func BenchmarkYears(b *testing.B) {
    for i := 0; i < b.N; i ++ {
        Years(100);
    }
}

func BenchmarkYearsTwo(b *testing.B) {
    for i := 0; i < b.N; i ++ {
        YearsTwo(200);
    }
}