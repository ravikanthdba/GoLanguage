/*

We are expeectng a pointeer receiver, but we are passing a non pointer value.
This would not work.

*/

package main

import (
	"fmt"
)

type square struct {
	length int;
	breadth int;
}

func (c *square) area() int {
    return c.length * c.breadth;
}

type shape interface {
	area() int;
}

func info(s *shape) { // This is accepting a pointer, hence this would not work, and we are thrown an error message.
	fmt.Printf("The area of a square is %d\n", s.area());
}

func main() {

	s1 := square{
		length: 10,
		breadth: 20,
	}

	fmt.Println(s1);
	fmt.Println(&s1);
	info(s1); // This is passing a value

}