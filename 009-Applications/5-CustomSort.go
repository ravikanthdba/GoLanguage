package main

import (
	"fmt"
	// "encoding/json"
	"sort"
)

type person struct {
	First string
	Last  string
	Age   float32
}

// func (p person) String() string {
// 	return fmt.Sprintf("%s %s : %f", p.First, p.Last, p.Age);
// }

type byAge []person
type byFirst []person

func (b byAge) Len() int {
	return len(b)
}

func (b byAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

func (b byFirst) Len() int {
	return len(b)
}

func (b byFirst) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byFirst) Less(i, j int) bool {
	return b[i].First < b[j].First
}

func main() {

	p1 := person{
		First: "Nagabhushanam",
		Last:  "Garimella",
		Age:   65,
	}

	p2 := person{
		First: "Swarna Latha",
		Last:  "Pratoori",
		Age:   55,
	}

	p3 := person{
		First: "Ravikanth",
		Last:  "Garmiella",
		Age:   31,
	}

	p4 := person{
		First: "Ratna Bhargavi",
		Last:  "Bharatula",
		Age:   27,
	}

	p5 := person{
		First: "Viraj Nandan",
		Last:  "Garimella",
		Age:   0.8,
	}

	p6 := person{
		First: "Kishore",
		Last:  "Devarakonda",
		Age:   34,
	}

	p7 := person{
		First: "Divya",
		Last:  "Garimella",
		Age:   27,
	}

	p8 := person{
		First: "Atharv",
		Last:  "Devarakonda",
		Age:   1,
	}

	peopleByAge := []person{p1, p2, p3, p4, p5, p6, p7, p8}

	fmt.Println("---------------- SORT BY AGE ----------------------")
	fmt.Println(peopleByAge)
	sort.Sort(byAge(peopleByAge))
	fmt.Println("-----------------------------")
	fmt.Println(peopleByAge)

	fmt.Println("\n\n")
	peopleByName := []person{p1, p2, p3, p4, p5, p6, p7, p8}
	fmt.Println("---------------- SORT BY NAME ----------------------")
	fmt.Println(peopleByName)
	sort.Sort(byFirst(peopleByName))
	fmt.Println("-----------------------------")
	fmt.Println(peopleByName)

}
