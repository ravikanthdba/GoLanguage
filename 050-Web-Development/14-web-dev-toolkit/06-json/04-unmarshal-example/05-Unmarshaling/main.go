package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Firstname  string `json:"Firstname"`
	LastName   string `json:"LastName"`
	Age        int    `json:"Age"`
	IsLinkedin bool   `json:"isLinkedin"`
}

func main() {
	var data []Data
	dataReceived := `[{"Firstname":"Ravikanth","LastName":"Garimella","Age":32,"isLinkedin":true},{"Firstname":"Nagabhushanam","LastName":"Garimella","Age":64,"isLinkedin":false},{"Firstname":"Swarna Latha","LastName":"Garimella","Age":59,"isLinkedin":false}]`
	err := json.Unmarshal([]byte(dataReceived), &data)
	if err != nil {
		fmt.Println(err)
	}
	for _, value := range data {
		fmt.Println(value)
	}
}
