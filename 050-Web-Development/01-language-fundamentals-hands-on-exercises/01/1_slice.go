/* composite literal - slice literal
composite literal is to have "type {}"
*/

package main

import "fmt"

func main() {
	var slice = []int{4, 5, 6, 7}
	for _, value := range slice {
		fmt.Println(value)
	}
}
