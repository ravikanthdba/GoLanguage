/* by mentioning fallthrough, when a case is tru and executed, automatically, the next case would also be executed.
In the below case, "Ravi", is passed and the case succeeds, it prints "It's Ravikanth" and since there is a fallthrough, even "You've entered Medi" also gets printed even if the case is true or not. Also "What's up Jenny" gets printed, since "Medi" case also has a fallthrough.
If I comment the fallthrough in Medi's case, then "What's up Jenny" doesn't get printed.
*/

package main

import (
	"fmt"
)

func main() {
	switch "Ravi" {
	case "Ravi":
		fmt.Println("It's Ravikanth")
		fallthrough
	case "Medi":
		fmt.Println("You've entered Medi")
		// fallthrough;
	case "Jenny":
		fmt.Println("What's up Jenny")
	default:
		fmt.Println("Do you have Friends")
	}
}
