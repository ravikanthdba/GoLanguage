package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type student struct {
	fname string
	lname string
	age int
}

func (s student) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, s.fname + " " + s.lname + " of age " + strconv.Itoa(s.age))
}

func main() {
	var s1 = student{
		fname: "Ravikanth",
		lname: "Garimella",
		age: 32,
	}

	http.ListenAndServe(":8080", s1)
}
