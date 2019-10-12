// recover halts the panic. It doesn't print the stack trace like the program in "panic.go"

/* The following are the sequence of steps that happen:

1) Main calls the PArty() function
2) In the Party(), 1st statement will be executed at the last as there is a defer.
3) In the Party(), 2nd statement will be executed at the second last as there is a defer.
4) First it prints "Hello".
5) Panic is called now.
6) Since panic is called, next step "Nice Weather!!" doesn't get printed.
7) colldown() function is called now.
8) recover() - halts the panic and stores under panicValue
9) panicValue is now printed, with the contents of whats printed in panic - without stacktrace
10) "Goodbye" - is now printed

*/

package main

import (
	"fmt"
)

func cooldown() {
	panicValue := recover()
	fmt.Println(panicValue)
}

func party() {
	defer fmt.Println("Goodbye")
	defer cooldown()

	fmt.Println("Hello")
	panic("I need to get out of here!!!")
	fmt.Println("Nice Weather!!")
}

func main() {
	party()
}

// Exercise: https://play.golang.org/p/tPNhOq3WADK