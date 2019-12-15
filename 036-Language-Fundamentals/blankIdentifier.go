/* If you declare a variable, and not use it, then its an error.
Blank identifier lets you handle that.
It is denoted by "_"
*/
package main

import (
	"fmt"
)

func main() {
	a := 0
	/* b := 1; This is failing with the below exception:
	   ╰─○ go run blankIdentifier.go
	       # command-line-arguments
	       ./blankIdentifier.go:13:5: b declared and not used
	       ╭─rgarimel at rgarimel-mn3 in ~/Documents/Programming/GoLang/goworkspace/src/github.com/ravikanthdba/GoLanguage/036-Language Fundamentals on master✘✘✘ using 19-11-05 - 18:49:27╰─○
	*/
	_ = 1
	fmt.Println(a)
}
