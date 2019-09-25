package main

import (
    "fmt"
)

func contains(a [][]int, b []int) bool {
	for i := 0; i < len(a); i ++ {
		if (a[i] == b) {
			return true
		}
	}
	return false
}


func birthday(s []int, d int, m int) {
    for i := 0; i < len(s); i ++ {
        for j := 0; j < len(s); j ++ {
            if (i != j) {
            	for count := 0; count < m; count ++ {
            		if (s[i] + s[j] == d) {
            			var x = []int{s[i], s[j]}
            		}
                }
            }
        }
        fmt.Println("---")
    }
}

func main() {
    var s = []int{2, 2, 1, 3, 2}
    var d, m int = 4, 2
    // var s = []int{1, 4}
    // var d, m int = 4, 1
    birthday(s, d, m)
}
