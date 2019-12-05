/* We cannot repeat the same conditions of case in 2 case statements. Example: Ravi cannot be in case 1 and case 2. It gives the following:
   ./switch-fallthrough.go:30:9: duplicate case "Jenny" in switch
    previous case at ./switch-fallthrough.go:28:22
*/

package main

import (
	"fmt"
)

func main() {
	switch "Jenny" {
	case "Ravi", "Jenny":
		fmt.Println("Hi Ravi or Jenny")
	case "Smith", "Raj":
		fmt.Println("Hi Smith or Raj")
	case "Rahul", "Preethi":
		fmt.Println("Hi Rahul or Preethi")
	}
}
