/*
   Key value pairs.
   can have nested Key value pairs.
   dictionary in python
   key - word; value - definition
   unordered
   uninitialized map is nil
   indexed by unique keys
   hash tables
*/

package main

import (
	"fmt"
)

func main() {
	c := make(map[int]string)
	c[1] = "a"
	c[2] = "b"
	fmt.Println(c)
	fmt.Println(&c)
	fmt.Printf("%T\n", c)
}
