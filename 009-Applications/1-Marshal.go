/*

Takes a struct and converets into json

*/

package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First string
	Last string
	Age int
}

func main() {

    p1 := person{
    	First: "Ravi",
    	Last: "Garimella",
    	Age: 31,
    }

    p2 := person{
    	First: "Viraj",
    	Last: "Nandan",
    	Age: 1,
    }

    people := []person{p1, p2}

    fmt.Println("The person p1 is ", p1)
    fmt.Print("The format of p1 is %T", p1)


    fmt.Println("\n\n")

    fmt.Println("The person p2 is ", p2)
    fmt.Print("The format of p2 is %T", p2)

    fmt.Println("\n\n")

    fmt.Println("The slice of person is people and the value of people is: ", people);
    fmt.Println("The format of people is %T", people);

    fmt.Println("\n\n")

    bs, _ := json.Marshal(people)
    fmt.Println("Converting into json format:", string(bs))
    fmt.Println("The format of bs is %T", bs); // Slice of bytw. It contains the ASC values of all the alphabets / numbers / symbols etc.

}