package main

import (
	"fmt"
	// "strings"
	"strconv"
)

func timeConversion(s string) string {
	var hour string
	getAMorPM := s[len(s) - 2 : len(s)]
	if (getAMorPM == "PM") { 
		var hourint int64
		hourint, err := strconv.ParseInt(s[0:2], 10, 64)
		if err != nil {
			fmt.Println(err)
		}
		return hourint
	} else {
		hour = s
		return hour
	}
}

func main() {
	time := "07:05:45PM"
	fmt.Println(timeConversion(time))
}