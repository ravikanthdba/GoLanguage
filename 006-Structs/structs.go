/*

   Data Structures - which composes values of different data type
   Aggregate together - values of different types
   Also called COMPOSITE DATA TYPE / COMPLEX DATA TYPES

*/

package main

import (
	"fmt"
)

type person struct {
	first_name string
	last_name  string
	age        int
}

func main() {

	// Example 1 - Create a person of type struct, with First_name, Last_name and age in it. Try to access the variables

	fmt.Println("Example 1 - Create a person of type struct, with First_name, Last_name and age in it. Try to access the variables")
	p1 := person{"Ravikanth", "Garimella", 31}

	p2 := person{"James", "Bond", 37}

	fmt.Println(p1)
	fmt.Println(p2)

	// The values in struct are accessed by '.' operator
	fmt.Println(p1.first_name, p1.last_name, p1.age)
	fmt.Println(p2.first_name, p2.last_name, p2.age)

	/* Example 2 - Embedded Structs

	   We can use a struct in another struct, and this is called as the embedded struct.
	   In the below example, we create a struct called secret agent, and we use person, in secret agent


	   Example Syntax:

	   type struct1 struct {
	       field datatype
	       field datatype
	       field datatype
	   }

	   type struct2 struct {
	       struct1
	       field datatype
	   }

	   While assigning, the syntax has to be:

	   st2 := struct2 {
	       struct1: struct1 {
	           field: value,
	           field: value,
	           field: value,
	       },
	       field: value,
	   }


	   Accessing:

	   fmt.Println(struct2.field)
	*/

	fmt.Println("\n\n")
	fmt.Println("Example 2 - Embedded Structs")
	type secretAgent struct {
		person
		license_to_kill bool
	}

	sa1 := secretAgent{
		person: person{
			first_name: "Ravikanth",
			last_name:  "Garimella",
			age:        32,
		},
		license_to_kill: false,
	}

	sa2 := secretAgent{
		person: person{
			first_name: "James",
			last_name:  "Bond",
			age:        37,
		},
		license_to_kill: true,
	}

	fmt.Println(sa1.person.first_name, sa1.person.last_name, sa1.person.age, sa1.license_to_kill)
	fmt.Println(sa2.person.first_name, sa2.person.last_name, sa2.person.age, sa2.license_to_kill)

	// If in the secret agent struct doesn't have an explicit first_name variable, then we can use it as sa1.first_name, instead of sa1.person1.first_name

	fmt.Println(sa1.first_name, sa1.last_name, sa1.age, sa1.license_to_kill)
	fmt.Println(sa2.first_name, sa2.last_name, sa2.age, sa2.license_to_kill)

	/* Anonymous structs

	   - Reduces Code pollution
	   - If a struct needs to be used at one single location, then we use an anonymous struct.
	   - In the above example2, it is not an anonymous struct, as the struct has a name, and the name is "person"

	   The below example shows: anonymous struct
	*/

	fmt.Println("\n\n")
	fmt.Println("Example 3 - Anonymous Structs")
	p3 := struct {
		father_first_name string
		father_last_name  string
		mother_first_name string
		mother_last_name  string
		father_age        int
		mother_age        int
	}{
		father_first_name: "Nagabhushanam",
		father_last_name:  "Garimella",
		father_age:        65,
		mother_first_name: "Swarna Latha",
		mother_last_name:  "Garimella",
		mother_age:        60,
	}

	fmt.Println(p3)

}
