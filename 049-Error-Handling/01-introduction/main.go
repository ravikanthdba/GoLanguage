/* In case of error, the error is printed to the console */

package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("File doesn't exist: ", err)
	}

	for i := 0; i < 5; i ++ {
		fmt.Println(i)
	}
}