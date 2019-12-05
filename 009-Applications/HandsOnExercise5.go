/*
Starting with this code, sort the []user by
age
last
Also sort each []string “Sayings” for each user
print everything in a way that is pleasant
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type byAge []user

func (u byAge) Len() int {
	return len(u)
}

func (u byAge) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u byAge) Less(i, j int) bool {
	return u[i].Age < u[j].Age
}

type byString []user

func (u byString) Len() int {
	return len(u)
}

func (u byString) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u byString) Less(i, j int) bool {
	return u[i].Last < u[j].Last
}

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	fmt.Println("-------- Sort by Age -----------")
	usersByAge := []user{u1, u2, u3}

	fmt.Println("Unsorted")
	for _, value := range usersByAge {
		fmt.Printf("First   :%s\n", value.First)
		fmt.Printf("Last    :%s\n", value.Last)
		fmt.Printf("Age     :%d\n", value.Age)
		fmt.Printf("Sayings :%s\n", value.Sayings)
		fmt.Println("\n\n")
	}
	fmt.Println("\n\n")
	sort.Sort(byString(usersByAge))
	fmt.Println("Sorted")
	for _, value := range usersByAge {
		fmt.Printf("First   :%s\n", value.First)
		fmt.Printf("Last    :%s\n", value.Last)
		fmt.Printf("Age     :%d\n", value.Age)
		fmt.Printf("Sayings :%s\n", value.Sayings)
		fmt.Println("\n\n")
	}

	fmt.Println("-------- Sort by Last Name -----------")

	usersByString := []user{u1, u2, u3}

	fmt.Println("Unsorted")
	for _, value := range usersByString {
		fmt.Printf("First   :%s\n", value.First)
		fmt.Printf("Last    :%s\n", value.Last)
		fmt.Printf("Age     :%d\n", value.Age)
		fmt.Printf("Sayings :%s\n", value.Sayings)
		fmt.Println("\n\n")
	}
	fmt.Println("\n\n")
	sort.Sort(byString(usersByString))
	fmt.Println("Sorted")
	for _, value := range usersByString {
		fmt.Printf("First   :%s\n", value.First)
		fmt.Printf("Last    :%s\n", value.Last)
		fmt.Printf("Age     :%d\n", value.Age)
		fmt.Printf("Sayings :%s\n", value.Sayings)
		fmt.Println("\n\n")
	}

	fmt.Println("-------- Sort by Sayings -----------")

	usersBySayings := []user{u1, u2, u3}
	for _, value := range usersBySayings {
		fmt.Println("Before Sorting, the value is:")
		fmt.Println("First Name: ", value.First)
		fmt.Println("Last  Name: ", value.Last)
		fmt.Println("Age       : ", value.Age)
		fmt.Println("Saying    : ", value.Sayings)
		fmt.Println(value)
		sort.Strings(value.Sayings)
		fmt.Println("After Sorting, the value is:")
		fmt.Println("First Name: ", value.First)
		fmt.Println("Last  Name: ", value.Last)
		fmt.Println("Age       : ", value.Age)
		fmt.Println("Saying    : ", value.Sayings)
	}

}
