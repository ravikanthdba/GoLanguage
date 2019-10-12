package main

import (
	"fmt"
	"strconv"
	"log"
)

func variableAssign() int {
	return 10
}

func multiReturnValues() (int, float32, string) {
	return 4, 0.123, "Ravikanth"
}

func main() {
	myVariable := variableAssign()
	fmt.Println(myVariable)
	var1, var2, var3 := multiReturnValues()
	fmt.Printf("var1 = %d, var2 = %f, var3 = %s\n", var1, var2, var3)



	fmt.Print("\n\n")

	flag, err := strconv.ParseBool("true")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(flag)


	flag, err = strconv.ParseBool("dummy")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(flag)
}