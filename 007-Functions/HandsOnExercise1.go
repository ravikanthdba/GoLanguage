/*


Hands on exercise
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results


*/

package main

import (
	"fmt"
)

func main() {

	n1 := foo();
	n2, n3 := bar();

	fmt.Println(n1);
	fmt.Println(n2);
	fmt.Println(n3);

}


func foo() int {
	return 42;
}

func bar() (int, string) {
	return 45, "Ravikanth";
}