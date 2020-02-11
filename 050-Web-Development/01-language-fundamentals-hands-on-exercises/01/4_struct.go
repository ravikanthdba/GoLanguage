/* composite literal; struct literal */

package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	p1 := person{
		firstname: "Ravikanth",
		lastname:  "Garimella",
	}

	fmt.Println("struct is now: ", p1)
}
