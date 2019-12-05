package main

import (
	"fmt"
)

func getDiagSum(matrix [][]int32) int32 {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			if i == j {
				sum += matrix[i][j]
			}
		}
	}
	return int32(sum)
}

func diagonalDifference(matrix [][]int32) int32 {
	getRegularDiagSum := getDiagSum(matrix)

	reversematrix := [][]int{}
	for i := 0; i <= len(matrix)-1; i++ {
		tempmatrix := []int{}
		for j := len(matrix) - 1; j >= 0; j-- {
			tempmatrix = append(tempmatrix, matrix[i][j])
		}
		reversematrix = append(reversematrix, tempmatrix)
	}

	getReverseDiagSum := getDiagSum(reversematrix)

	if (getRegularDiagSum - getReverseDiagSum) > 0 {
		return (getRegularDiagSum - getReverseDiagSum)
	} else {
		return (getReverseDiagSum - getRegularDiagSum)
	}

}

func main() {
	var matrix = [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(diagonalDifference(matrix))
}
