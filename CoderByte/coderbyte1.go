package main

import (
	"fmt"
	"unicode"
)

func LetterChanges(str string) string {
	var output string = ""

	for _, letter := range str {
		if unicode.IsLetter(letter) {
			if unicode.ToUpper(letter+1) == 'A' || unicode.ToUpper(letter+1) == 'E' || unicode.ToUpper(letter+1) == 'I' || unicode.ToUpper(letter+1) == 'O' || unicode.ToUpper(letter+1) == 'U' {
				output += string(unicode.ToUpper(letter + 1))
			} else {
				output += string(letter + 1)
			}
		} else {
			output += string(letter)
		}
	}
	return output
}

func main() {

	fmt.Println(LetterChanges("hello*3"))

}
