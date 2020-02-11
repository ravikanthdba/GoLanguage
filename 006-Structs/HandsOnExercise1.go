/*

Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

*/

package main

import (
	"fmt"
)

type person struct {
	first_name                 string
	last_name                  string
	favourite_ice_cream_flavor []string
}

func main() {

	p1 := person{first_name: "Bhargavi",
		last_name:                  "Bharatula",
		favourite_ice_cream_flavor: []string{"chocolate", "butterscotch", "mango"},
	}
	p2 := person{first_name: "Divya",
		last_name:                  "Garimella",
		favourite_ice_cream_flavor: []string{"vanilla", "mango", "strawberry"},
	}

	fmt.Printf("Name : %s %s\n", p1.first_name, p1.last_name)
	fmt.Println("Favourite Icecream flavours are: ")
	for icecreamflavour := 0; icecreamflavour < len(p1.favourite_ice_cream_flavor); icecreamflavour += 1 {
		fmt.Printf("%d.   \t %s\n", icecreamflavour, p1.favourite_ice_cream_flavor[icecreamflavour])
	}

	fmt.Println("\n\n")

	fmt.Printf("Name: %s %s\n", p2.first_name, p2.last_name)
	fmt.Println("Favourite Icecream flavours are: ")
	for icecreamflavour := 0; icecreamflavour < len(p2.favourite_ice_cream_flavor); icecreamflavour += 1 {
		fmt.Printf("%d.   \t %s\n", icecreamflavour, p2.favourite_ice_cream_flavor[icecreamflavour])
	}

	fmt.Println("\n\n")
	fmt.Println("Common Icecream Flavours for both of them are:")
	for p1icecreamflavour := 0; p1icecreamflavour < len(p1.favourite_ice_cream_flavor); p1icecreamflavour += 1 {
		for p2icecreamflavour := 0; p2icecreamflavour < len(p2.favourite_ice_cream_flavor); p2icecreamflavour += 1 {
			if p1.favourite_ice_cream_flavor[p1icecreamflavour] == p2.favourite_ice_cream_flavor[p2icecreamflavour] {
				fmt.Println(p1.favourite_ice_cream_flavor[p1icecreamflavour])
			}
		}
	}
}
