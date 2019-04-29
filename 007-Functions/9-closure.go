    /*

    Defines the scope of the variable

    */

    package main

    import (
        "fmt"
    )

    var x int; // The scope of x is the entire package. It can be ecalled anywhere in the package.

    func main() {
        var y int = 42;
        fmt.Printf("From main: x = %d\n", x);
        foo();

        fmt.Printf("From main: y = %d\n", y);


        {
        	var z int = 100;
        	fmt.Printf("A code block under main: z = %d\n", z);
        }

        // fmt.Printf("From main: z = %d\n", z); // this would fail, as the scope of z is limited to that cpde block alone.



        one := incrementor()
        two := incrementor()

        fmt.Println("The value of one is ", one());
        fmt.Println("The value of one is ", one());
        fmt.Println("The value of one is ", one());
        fmt.Println("The value of one is ", one());
        fmt.Println("The value of one is ", one());
        fmt.Println("The value of one is ", one());
        fmt.Println("The value of two is ", two());
        fmt.Println("The value of two is ", two());
    }

    func foo() {
        x ++;
        fmt.Printf("From Foo: x = %d\n", x);


        // fmt.Printf("From foo: y = %d\n", y); // This would fail, as the scope of y is limited to main functioni
    }

    func incrementor() func() int {
    	var a int;
    	return func() int {
    		a ++;
    		return a;
    	}
    }