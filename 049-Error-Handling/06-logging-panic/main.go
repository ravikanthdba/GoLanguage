/* If you need more stacktrace then you do a panic. It returns exit status 2 */

package main

import (
	"os"
)

func main() {
	_, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
}
