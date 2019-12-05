package main

import (
	"fmt"
)

func main() {
	var x = [2][2]int{{1, 2}, {3, 4}}
	var y = [2][2]int{{9, 10}, {11, 12}}

	for i := 0; i < len(x); i++ {
		for j := 0; j < len(x[i]); j++ {
			fmt.Println(x[i][j] + y[i][j])
		}
	}

}
