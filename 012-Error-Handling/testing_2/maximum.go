package main


func maximum(numbers []int) int {
    var high int = 0;
    for value := range numbers {
        if numbers[value] > high {
            high = numbers[value];
        } else {
            continue;
        }
    }
    return high;
}