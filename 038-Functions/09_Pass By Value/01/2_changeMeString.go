package main

import (
	"fmt"
)

func main() {
	var x string = "Hello"
	fmt.Println("Main: Before changing, the value of hello is: ", x)
	changeMeString(&x)
	fmt.Println("Main: After changing, the value of hello is: ", x)
}

func changeMeString(z *string) {
	fmt.Println("changeMeString: Before changing, the value of z is: ", z)
	*z = "world"
	fmt.Println("changeMeString: After changing, the value of z is: ", z)
}
