package main

import (
	"fmt"
)

func main() {
	// var x map[string]string;

	/* The below doesn't work as the map defined "x" is a nil map */

	// x["Todd"] = "Good Morning!!";
	// x = append(x, "Todd");
	// fmt.Println(x);

	y := map[string]string{} // Composite literal
	y["Jane"] = "Good Afternoon!!"
	y["Tom"] = "Good Morning Tom!!"
	fmt.Println(y)

	// Fetch:
	fmt.Println(y["Jane"], y["Tom"])

	// Length of map:
	fmt.Println("y is of length: ", len(y))

	// deleting a map:
	delete(y, "Jane")
	fmt.Println(y)
}
