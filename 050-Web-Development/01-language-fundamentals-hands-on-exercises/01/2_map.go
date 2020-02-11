/* composite literal; map literal */

package main

import "fmt"

func main() {
	var dict = map[string]int{"Ravikanth": 1, "Garimella": 2}
	for _, value := range dict {
		fmt.Println(value)
	}

	for key, _ := range dict {
		fmt.Println(key)
	}
}
