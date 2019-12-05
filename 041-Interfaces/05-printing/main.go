/*
	Overriding the regular fmt.Println method
*/

package main

import "fmt"

type Person struct {
	First string
	Last  string
}

func (p Person) String() string {
	return "YAYAYAYYAYA!!!!" + p.First + " " + p.Last
}

func main() {
	p1 := Person{
		First: "Ravikanth",
		Last:  "Garimella",
	}

	fmt.Println(p1)
}
