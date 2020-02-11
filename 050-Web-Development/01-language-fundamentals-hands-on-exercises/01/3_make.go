/* map using make */

package main

import "fmt"

func main() {
	dict := make(map[string]int)
	fmt.Println(dict)

	dict["Ravikanth"] = 1
	dict["Garimella"] = 2

	fmt.Println(dict)

	for key, value := range dict {
		fmt.Println(key, value)
	}

	dict["Garimella"] = 4
	for key, value := range dict {
		fmt.Println(key, value)
	}
}
