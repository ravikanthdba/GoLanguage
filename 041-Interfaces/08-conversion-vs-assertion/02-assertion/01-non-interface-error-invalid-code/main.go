/* Since "name" variable is not of type interface, the below code fails */

package main

import (
	"fmt"
)

func main() {
	name := "Ravikanth"
	str, ok := name.(string)
	if ok {
		fmt.Println(str)
	} else {
		fmt.Println("Value not string")
	}

}
