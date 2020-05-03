package main

import (
	"fmt"
)

func gradingStudents(grades []int32) []int32 {
	var finalgrade []int32

	for i := 0; i < len(grades); i++ {
		var grade int32
		grade = grades[i]
		if grades[i] >= 38 {

			for (grades[i]%int32(5) != 0) && (grades[i]%int32(5) != 5) {
				grades[i] += 1
			}

			if grades[i]-grade < 3 {
				finalgrade = append(finalgrade, grades[i])
			} else {
				finalgrade = append(finalgrade, grade)
			}
		} else {
			finalgrade = append(finalgrade, grade)
		}
	}

	return finalgrade
}

func main() {
	var grades = [4]int32{73, 67, 38, 33}
	fmt.Println(gradingStudents(grades))
}
