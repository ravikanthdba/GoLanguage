package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "123456789"
	x := "10.123456789"
	stringConverted, _ := strconv.ParseInt(s, 10, 32)
	fmt.Println(stringConverted)
	stringConverted, _ = strconv.ParseInt(x, 10, 32)
	fmt.Println(stringConverted)
	stringConvertedToFloat, _ := strconv.ParseFloat(x, 64)
	fmt.Println(stringConvertedToFloat)
}
