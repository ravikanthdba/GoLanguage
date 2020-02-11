/* Convert a Json to slice of bytes */

package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {

	s := `[{"First":"Ravi","Last":"Garimella","Age":31},{"First":"Viraj","Last":"Nandan","Age":1}]`
	bs := []byte(s)

	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("All people:", people)
	for i, v := range people {
		fmt.Println("Person number: ", i)
		fmt.Println("First Name:    ", v.First)
		fmt.Println("Last Name:     ", v.Last)
		fmt.Println("Age:           ", v.Age)
		fmt.Println("\n\n")
	}

}
