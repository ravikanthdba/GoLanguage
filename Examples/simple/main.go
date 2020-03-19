package main

import (
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
	"encoding/json"
)

type Oncall []struct {
	FullName string   `json:"full_name"`
	Start    int      `json:"start"`
	End      int      `json:"end"`
	User     string   `json:"user"`
	Team     string   `json:"team"`
	Role     string   `json:"role"`
}

func main() {
	resp, err := http.Get("https://oncall.prod.linkedin.com/api/v0/teams/search-sre/oncall")
	if err != nil {
		log.Fatalln("Unable to hit the end point: , erroring out with the error: ", err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Error converting the data: ", err)
	}

	var oncall Oncall
	err = json.Unmarshal(data, &oncall)
	if err != nil {
		log.Fatalln("Error converting the JSON to struct: ", err)
	}

	for _, oncallPerson := range oncall {
		if oncallPerson.Role == "primary" {
				fmt.Println(oncallPerson.FullName, oncallPerson.Role)
			}
		}
}