/*
Starting with this code, encode to JSON the []user sending the results to Stdout. 
Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})
*/

package main

import (
	"fmt"
	"encoding/json"
	"os"
)

type person struct {
	First string
	Last string
	Age float32
}

func main() {

	p1 := person {
		First: "Nagabhushanam",
		Last: "Garimella",
		Age: 65,
	}


	p2 := person {
		First: "Swarna Latha",
		Last: "Pratoori",
		Age: 55,
	}

    p3 := person {
		First: "Ravikanth",
		Last: "Garmiella",
		Age: 31,
	}

	p4 := person {
		First: "Ratna Bhargavi",
		Last: "Bharatula",
		Age: 27,
	}

	p5 := person {
		First: "Viraj Nandan",
		Last: "Garimella",
		Age: 0.8,
	}

	p6 := person {
		First: "Kishore",
		Last: "Devarakonda",
		Age: 34,
	}

	p7 := person {
		First: "Divya",
		Last: "Garimella",
		Age: 27,
	}

	p8 := person {
		First: "Atharv",
		Last: "Devarakonda",
		Age: 1,
	}

	people := []person{p1, p2, p3, p4, p5, p6, p7, p8};

	err := json.NewEncoder(os.Stdout).Encode(people)
	if err != nil {
		fmt.Println(err);
	}
	
}