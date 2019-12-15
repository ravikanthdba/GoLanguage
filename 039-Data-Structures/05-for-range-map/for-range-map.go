package main

import (
	"fmt"
)

func main() {
	var mapVariable = map[int]string{
		1: "Good Morning",
		2: "Subhodayam",
		3: "Aadab hyderabad",
	}

	var mapVariable2 = map[string]map[string]string{
		"Ravikanth": {
			"Language": "Telugu",
		},
		"Swarna": {
			"Language": "Telugu",
		},
	}

	for id, variable := range mapVariable {
		fmt.Println(id, variable)
	}

	for id2, variable2 := range mapVariable2 {
		fmt.Println(id2, variable2, variable2["Language"])
	}
}
