// This function takes a string and reverses it and prints it.

package myreversefunction

func Reverse(s string) string {
	var reversed_string string = ""

	for str := len(s) - 1; str >= 0; str-- {
		reversed_string += string(s[str])
	}

	return reversed_string
}
