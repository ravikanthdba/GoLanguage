/*

There are 4 categories:

- NON POINTER RECEIVERS and NON POINTER VALUES
- NON POINTER RECEIVERS and POINTER VALUES
- POINTER RECEIVRES and NON POINTER VALUES
- POINTEER RECEIVRES and POINTER VALUES

A pointer receiver works with both nonpointer and pointer values
A non pointer receiver works with only non pointer values

*/

package main

import (
    "fmt"
)

const pi = 3.14157;

type circle struct {
    radius float32;
}


type square struct {
    length float32;
    breadth float32;
}

func (c circle) area() { // Info also has a non-pointer receiver
    fmt.Println(pi * c.radius * c.radius);
}

func (s square) area() {
    fmt.Println(s.length * s.breadth);
}

type shape interface {
    area();
}

func info(s shape) {
    s.area();
}

func main() {

    c1 := circle{
            radius: 10,
        }

    info(c1); // In this case, we pass non pointer value to the info function

    s1 := square{
        length: 10,
        breadth: 10,
    }

    info(s1)

}