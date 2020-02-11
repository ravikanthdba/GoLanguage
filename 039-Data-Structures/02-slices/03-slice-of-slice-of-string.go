package main

import (
	"fmt"
)

func main() {
	records := make([][]string, 0)
	student1 := make([]string, 4)
	student1[0] = "Ravi"
	student1[1] = "Garimella"
	student1[2] = "90"
	student1[3] = "80"

	records = append(records, student1)

	student2 := make([]string, 4)
	student2[0] = "xyz"
	student2[1] = "yzx"
	student2[2] = "95"
	student2[3] = "85"

	records = append(records, student2)

	fmt.Println(records)
}
