/*

Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

*/

package main

import (
	"fmt"
)

type person struct {
	first_name          string
	last_name           string
	favourite_ice_cream []string
}

func main() {

	p1 := person{
		first_name:          "Bhargavi",
		last_name:           "Bharatula",
		favourite_ice_cream: []string{"chocolate", "butterscotch", "vanilla"},
	}

	p2 := person{
		first_name:          "Divya",
		last_name:           "Garimella",
		favourite_ice_cream: []string{"chikku", "strawberry", "chocolate"},
	}

	m := map[string]person{
		p1.last_name: p1,
		p2.last_name: p2,
	}

	for _, value := range m {
		fmt.Println(value.first_name)
		fmt.Println(value.last_name)

		fmt.Println("\n\n")
		fmt.Println("Favourite Icecream Flavours are:")
		for _, v := range value.favourite_ice_cream {
			fmt.Println(v)
		}
		fmt.Println("---------------------------------------")
	}

}
