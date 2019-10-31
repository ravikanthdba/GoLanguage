package main

import (
    "fmt"
)

func contains(elementslist []bool, value bool) bool {
    for _, element := range elementslist {
        if element == value {
            return true
        } 
    }
    return false
}

func getTotalX(a []int, b []int) {
    
    var n int = 1
    var final []int
        
    for ((n < len(a) || n < len(b)) || n <= a[len(a) - 1] || n <= b[len(b) - 1]) {
        
        var result []bool

        for i := 0; i < len(a); i ++ {
            
            if (n % a[i] == 0) {
                result = append(result, true)
            } else {
                result = append(result, false)
            }
        }


        for i := 0; i < len(b); i++ {
            if (b[i] % n == 0) {
                result = append(result, true)
            } else {
                result = append(result, false)
            }
        }



        if (contains(result, false) != true) {
            final = append(final, n)
        }

        n += 1
    }

    fmt.Println(len(final))
    
}

func main() {
    var a = []int {2, 6}
    var b = [] int {24, 36}

    getTotalX(a, b)
}