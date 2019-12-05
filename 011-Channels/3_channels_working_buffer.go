package main

import (
	"fmt"
)

func main() {

	/* Sample: 1
	   1st LOC -  The below code can have 1 value in the buffer.
	   2nd LOC -  Hence, it was able to store 44 into "c"
	   3rd LOC -  Prints the channel from "c"
	*/
	fmt.Println("Sample: 1")
	sample1_channel := make(chan int, 1)
	sample1_channel <- 44
	fmt.Println(<-sample1_channel)

	fmt.Println("Sample: 2 - It wouldn't print anything. The logic is in the code, on why it doesn't print anything. Please check the code.")
	/* Sample 2:
	   The below code doesn't work, as we can store only one variable in the buffer. Here we are passing two elements into the buffer, and printing it only once.

	   sample2_channel := make(chan int, 1);
	   sample2_channel <- 44
	   sample2_channel <- 45
	   fmt.Println(<- sample2_channel)

	*/

	/* Sample 3:
	   LOC 1: Creates a channel with buffer 1
	   LOC 2: Adds 33 to sample3_channel, it maintains 33 in the buffer created in the channel.
	   LOC 3: Prints 33, and the buffer gets emptied.
	   LOC 4: Adds 44 to sample3_channel, it maintains 44 in the buffer created in the channel.
	   LOC 5: Prints 44 to sample3_channel and the buffer gets emptied again.

	   The below code works!!
	*/
	fmt.Println("Sample: 3")
	sample3_channel := make(chan int, 1)
	sample3_channel <- 33
	fmt.Println(<-sample3_channel)
	sample3_channel <- 44
	fmt.Println(<-sample3_channel)

	/*
	   Sample 4:
	   LOC 1: Creates a channel sample4_channel with 2 buffers.
	   LOC 2: Adds 33 to buffer.
	   LOC 3: Adds 44 to buffer.
	   LOC 4: Prints the channel once, which means since 33 is in the top of the stack, it will print 33, and then there will be 44 at the top of the stack.
	   LOC 5: Another print of the channel, it pops out 44 onto the screen and the channel becomes empty now.
	   LOC 6: Would not work, as the number of buffers are now 0
	*/
	fmt.Println("Sample: 4")
	sample4_channel := make(chan int, 2)
	sample4_channel <- 33
	sample4_channel <- 44
	fmt.Println(<-sample4_channel)
	fmt.Println(<-sample4_channel)
	// fmt.Println(<- sample4_channel);

	/* Below is a sample channel with strings */

	fmt.Println("Sample: 5")
	sample5_channel := make(chan string, 5)
	sample5_channel <- `a`
	sample5_channel <- `b`
	sample5_channel <- `c`
	sample5_channel <- `d`
	sample5_channel <- `2`

	fmt.Println(<-sample5_channel)
	fmt.Println(<-sample5_channel)
	fmt.Println(<-sample5_channel)
	fmt.Println(<-sample5_channel)
	fmt.Println(<-sample5_channel)
}
