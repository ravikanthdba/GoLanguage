/* panic cannot be captured in a log file, hence this code doesn't work */

package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	logfile, err := os.Create("logging.txt")
	if err != nil {
		fmt.Println("Unable to create the file")
	}
	log.SetOutput(logfile)

	_, err := os.Open("input.txt")
	if err != nil {

	}
}
