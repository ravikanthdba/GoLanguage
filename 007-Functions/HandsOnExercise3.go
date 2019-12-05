/*

Use the “defer” keyword to show that a deferred func runs after the func containing it exits.

*/

package main

import (
	"fmt"
)

func main() {
	first()
	second()
	defer third() // This should execute last, even after the second defer below
	fourth()
	defer fifth()
	sixth()

}

func first() {
	fmt.Println("Executing first")
}

func second() {
	fmt.Println("Executing second")
}

func third() {
	fmt.Println("Executing third")
}

func fourth() {
	fmt.Println("Executing fourth")
}

func fifth() {
	fmt.Println("Executing fifth")
}

func sixth() {
	fmt.Println("Executing sixth")
}
