// Entry point to the program
// Exit to the program is when it ecits the main function

/*

In the below program, the program starts with main function and goes in sequential order.
First it prints "GoLanguage is Fun!!"
next - it gors to the foo function
It then prints "I am in 'foo' function"
Next it prints "Printing in next line"
Then it hits the closed parenthesis of the main function. This is when the code would exit.
Hence, the code would end here.
*/

package main // 	Compulsory

import "fmt"

func main() {
	n, _ := fmt.Println("GoLanguage is Fun!!")
	fmt.Println("The number of bytes are ", n);
	foo()
	fmt.Println("Printing in next line")

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println("Even Number : ", i)
		} else {
			fmt.Println("Odd Number  : ", i)
		}
	}
}

func foo() {
	fmt.Println("I am in 'Foo' function")
}
