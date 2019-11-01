/*
    Start with https://play.golang.org/p/9a1IAWy5E6 Create a custom error message using “fmt.Errorf”.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	c := "string to be passed"

	bs, err := toJSON(p1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))

	bs, err = toJSON(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(bs))
}

// toJSON needs to return an error also
func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		errorMessage := fmt.Errorf("Customized error called")
		return nil, errorMessage
	}
	return bs, nil
}
