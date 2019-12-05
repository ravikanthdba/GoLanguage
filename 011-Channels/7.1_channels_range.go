/*
We can select from the channel, using the "range" query.
The "range" query loops over the channel, until a statement close(channel) is found, and prints from the channel.

In the below example: we have a channel called "queue", which maintains 5 elements in the buffer, which is of integer type. "chan int" - tells us, that we have created a channel with integer type.
We assign 5 elements to the channel by calling the "queue <- value", 5 times as we mentioned above, we can maintain 5 elements in the channel buffer.
We then call the close(queue) function, which tells, there will be no more messages propogated to the channel queue.

We then loop over the channel named "queue", and then print the value in the channel.
It works as a queue process - First In First out.

Hence when we print, its output is: 1, 2, 3, 4, 5
*/

package main

import (
	"fmt"
)

func main() {
	queue := make(chan int, 5)

	queue <- 1
	queue <- 2
	queue <- 3
	queue <- 4
	queue <- 5

	close(queue)

	for value := range queue {
		fmt.Println("Value is: ", value)
	}
}
