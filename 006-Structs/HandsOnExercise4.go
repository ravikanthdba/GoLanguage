/*

Create and use an anonymous struct.

*/

package main

import (
	"fmt"
)

func main() {

	p1 := struct {
		first_name         string
		last_name          string
		age                int
		favourite_icecream []string
	}{
		first_name:         "Ravikanth",
		last_name:          "Garimella",
		age:                32,
		favourite_icecream: []string{"Vanilla", "Chocolate", "butterscotch"},
	}

	fmt.Println(p1.first_name)
	fmt.Println(p1.last_name)
	fmt.Println(p1.age)
	for icecream := 0; icecream < len(p1.favourite_icecream); icecream++ {
		fmt.Println(p1.favourite_icecream[icecream])
	}
}
