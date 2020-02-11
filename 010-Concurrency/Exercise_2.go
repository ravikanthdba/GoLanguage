/*

This exercise will reinforce our understanding of method sets:
create a type person struct
attach a method speak to type person using a pointer receiver
*person
create a type human interface
to implicitly implement the interface, a human must have the speak method
create func “saySomething” 
have it take in a human as a parameter
have it call the speak method
show the following in your code 
you CAN pass a value of type *person into saySomething
you CANNOT pass a value of type person into saySomething

*/

package main

import (
	"fmt"
	"strconv"
)

type person struct {
	name string
	age int
}

func (p *person) speak() string {
	return (p.name  + " of age: " + strconv.Itoa(p.age) + " speaks!!")
}

func (p person) speakWithoutPointerReceiver() string {
	return p.name + " says hello, and the call for this function is without receiver pointer."
}

type human interface {
	speak() string
}

type humanWithoutPointerReceiver interface {
	speakWithoutPointerReceiver() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func saySomethingWithoutPointerReceiver(h humanWithoutPointerReceiver) {
	fmt.Println(h.speakWithoutPointerReceiver())
}

func main() {
	p1 := person{
		name: "Ravikanth",
		age: 32,
	}

	saySomething(&p1)
	// saySomething(p1) // does not work - as the speak takes pointer as a receiver. If the receiver is a pointer, then it takes only pointers as an input. If the receiver is not a pointer, then it takes pointer and normal value as input.
	fmt.Println("Receiver is not a pointer, but value passed is a pointer: ")
	saySomethingWithoutPointerReceiver(&p1)
	fmt.Println("Receiver is not a pointer, but value passed is also not a pointer: ")
	saySomethingWithoutPointerReceiver(p1)

}
