/* Remainder of each number when divided by 4. The range of numbers are : 10 to 100 */

package main


import (
	"fmt"
)

func main() {
	for number := 10; number <= 100; number ++ {
		fmt.Printf("The number %d when divided by 4 is : %d\n", number, (number % 4));
	}
}