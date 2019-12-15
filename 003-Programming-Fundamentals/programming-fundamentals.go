/*
    This program denotes the Numeric Types in Golang.

    There are couple of Integer datatypes in Golang:


    uint8 -   8 bit Integers - 0 to 255
    uint16 - 16 bit Integers - 0 to 65535
    uint32 - 32 bit - 0 to 2^32 - 1
    uint64 - 64 bit - 0 to 2^64 - 1


    int8 - 8 bit Integers - -128 to 127
    int16 - 16 bit Integers - -32768 to 32767
    int32 - 32 bit Integers - -2^32 to 2^32 - 1
    int64 - 64 bit Integers - -2^64 to 2^64 - 1


    float32 - 32bit float Integers
    float64 - 64bit float Integers


    complex64 - float32 real and Imaginary numbers
    complex128 - float64 real and Imaginary numbers


    byte - alias for uint8
    rune - alias for int32


    runtime - provides the system configuration details, where the Go Program is being run.

    GOOS   - Provides the Operating System Details, in case of MAC, it is "darwin"
    GOARCH - Provides the Processor Architecture, in the case of current system, it is AMD 64bit

    Strings are mutable, and we cannot change the string, once assigned to a variable.


    _ variable is created - when you are assigning a variable and not using it.
*/

package main

import (
	"fmt"
	"runtime"
)

var integer int
var decimal float32


const (
    	_ = iota
    	kb = 1 << (iota * 10)
    	mb = 1 << (iota * 10)
        gb = 1 << (iota * 10)
    )

func main() {
	var x bool
	fmt.Printf("The datatype of x is %T and the value of x is %t\n", x, x)

	x = true
	fmt.Printf("The datatype of x is %T and the value of x is %t\n", x, x)

	integer = 100
	decimal = 60.123

	fmt.Printf("The datatype of integer is %T and the value of integer is %d\n", integer, integer)
	fmt.Printf("The datatype of decimal is %T and the value of decimal is %f\n", decimal, decimal)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOMAXPROCS)


	var str string = "Hello World";
	fmt.Println(str);
	fmt.Printf("%T\n", str);

	bs := []byte(str);
	fmt.Println(bs);
	fmt.Printf("%T\n", bs);

	for i := 0; i < len(str); i ++ {
	    fmt.Printf("%#U\n", str[i]);
	}

	for i, v := range str {
	    fmt.Printf("At Index %d, the value of the Hexadecimal value is %#U, and the value is %c\n", i, v, str[i]);
	}



	var s string = "H";
	fmt.Println(s);


	cs := []byte(s);
	fmt.Println(cs);


	n := cs[0];
	fmt.Printf("%T\n", n);


    // Convert n to binary and hexadecimial case

    fmt.Printf("The binary value for %d is %b\n", n, n);
    fmt.Printf("The Hexadecimal value for %d is %#x\n", n, n);
    fmt.Printf("The Hexadecimal value for %d is %#X\n", n, n);



    /* Constantants

    - Const is the keyword in Go for constants

    */


    const a = 42;
    const b = 42.28;
    const c = "Go Programming Language";


	// Untyped constants

    const (
    	alpha = 99;
    	beta = 99.98;
    	gamma = "Golang"
    )

    fmt.Println(a);
    fmt.Println(b);
    fmt.Println(c);

    fmt.Printf("%T\n", a);
    fmt.Printf("%T\n", b);
    fmt.Printf("%T\n", c);

    fmt.Println(alpha);
    fmt.Println(beta);
    fmt.Println(gamma);

    fmt.Printf("%T\n", alpha);
    fmt.Printf("%T\n", beta);
    fmt.Printf("%T\n", gamma);



/*
	 Within a constant declaration, the pre-defined iota represents successive untyped integer.
	 In the below case, c0 = iota means c0 = 0; c1 = iota - means - value of c0 + 1 = 1 and so forth
	 Iota resets when const is defined again, in teh below scenario, c3 and c4 are defined again as constants.
	 If c3 and c4 - has been declared in the same const declaration like 
    const (
    	c0 = iota
    	c1 = iota
    	c2 = iota
    )

*/

    const (
    	c0 = iota
    	c1 = iota
    	c2 = iota
    )

    const (
    	c3 = iota
    	c4 = iota
    )

    fmt.Println(c0);
    fmt.Println(c1);
    fmt.Println(c2);
    fmt.Println(c3);
    fmt.Println(c4);


    /* Bit Shifting - moving the bit 1 by by x places
        << - left shift
        >> - Right shift

        x = 4 -> x = 100 (In Binary) -> 0100
        x << 1 -> 1000 - 8


        x = 4 -> x = 100 -> 0100
        x << 2 -> 10000 -> 32
    */


    var integer int = 4;
    binary := fmt.Sprintf("%b", integer);
    fmt.Println(binary);
    fmt.Println("Shifting bits to the left")
    left_shift := integer << 4;
    fmt.Printf("After four left shifting to the number - %d  the binary is  %b\n", left_shift, left_shift);
    right_shift := integer >> 1;
    fmt.Printf("After four right shifting to the number - %d  the binary is  %b\n", right_shift, right_shift);


    fmt.Printf("The binary for %d is %b\n", kb, kb);
    fmt.Printf("The binary for %d is %b\n", mb, mb);
    fmt.Printf("The binary for %d is %b\n", gb, gb);


    // fmt.Printf("The values are %b %b and %b\n", KB << 10, MB << 20, GB << 30);

    var year int;
    fmt.Println("Enter the year to find out whether it is a leap year or non leap year");
    fmt.Scanf("%d", &year);
    fmt.Printf("The year entered is %d\n", year);
    if (year % 4 == 0) {
    	fmt.Printf("%d is a Leap Year\n", year);
    } else {
    	fmt.Printf("%d is a Non Leap Year\n", year);
    }

    
    
}
