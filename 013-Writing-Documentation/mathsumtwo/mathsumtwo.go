// This function is for summing the numbers.

package mathsumtwo

func MathSumTwo(variables []int) int {
	var sum int = 0

	for values := range variables {
		sum += variables[values]
	}

	return sum
}
