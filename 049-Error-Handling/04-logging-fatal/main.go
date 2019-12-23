/* log.Fatal calls os.Exit(1), if an error occurs.
os.Exit(0) - is a success
os.Exit(1) - is a failure
*/

package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("error for the command: ", err)
	}
}
