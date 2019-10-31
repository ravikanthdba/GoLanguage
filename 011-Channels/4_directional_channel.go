/*
    The channels created in 3_channels_working_buffer.go and previous codes are bi-directional channels.
    The directional channels mean, we can either only send to the channel (or) receive from the channel.

    Send only channel is given by "chan <- int"
    Receive only channel is given by "<- chan int"

    A bi-directional channel can be changed to Send only channel (or) receive only channel, but vice-versa is not true.

    We cannot assign values to receive-only channel, and we cannot receive from send-only channel.
*/

package main

import (
    "fmt"
)

func main() {
    bidirectional := make(chan int, 1); // Bi-directional channel with buffer of 1
    bidirectional <- 44; // Sending events to bi-directional channel
    fmt.Println("The value of bidirectional channel is: ", <- bidirectional); // receiving events from bidirectional channel.
    fmt.Printf("Type of bidirectional is: %T \n", bidirectional)


    // unidirectional_send_only := make(chan <- int, 1); // uni-directional send only channel, with a buffer of 1
    // unidirectional_send_only <- 44;
    // fmt.Println("The value of unidirectional send only channel is:", <- unidirectional_send_only); // This code will not work, we are trying to receive from a send only channel.

    // unidirectional_receive_only := make(<- chan int, 1); // uni-directional channel receive only, with a buffer of 1
    // unidirectional_receive_only <- 44; // This code would not work, as we are trying to send an event to a receive only channel.
    // fmt.Println("The value of unidirectional_receive_only channel is: ", <- unidirectional_receive_only) // this code would not work as the above code is not working. If we comment the above code, and try re-running, since there is no value sent to the channel, receiving is not possible.

    unidirectional_send_only := make(chan <- int, 1);
    fmt.Printf("Type of unidirectional_send_only is: %T \n", unidirectional_send_only)

    unidirectional_receive_only := make(<- chan int, 1);
    fmt.Printf("Type of unidirectional_receive_only is: %T \n", unidirectional_receive_only)

   unidirectional_send_only = bidirectional
   unidirectional_receive_only = bidirectional

   fmt.Println("After assignment: ")
   unidirectional_send_only = make(chan <- int, 1);
   fmt.Printf("Type of unidirectional_send_only is: %T \n", unidirectional_send_only)

   unidirectional_receive_only = make(<- chan int, 1);
   fmt.Printf("Type of unidirectional_receive_only is: %T \n", unidirectional_receive_only)

}

