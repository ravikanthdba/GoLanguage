// This package is used for summation
package summation

// This function takes any number of input parameters and sums them up.
func Summation(n ...int) int {
    sum := 0;

    for _, value := range n {
        sum += value;
    }

    return sum;
}
