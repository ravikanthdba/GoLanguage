/*

Two parts to the chapter:

	- Loops - Number of times the command has to execute
	- Conditions - Execute next steps based on a Conditions
	- GoLamg does not have a while, it only has a For loop


	Structure:


	for init; condition; incrememnt {
		statement
	}


	- For Loop keeps on running until the condition evaluates to true. This is the similar case of a while loop.

	- Conditional Logic Operators:

	&&  -- and
	||  -- or
	 !  -- not
*/

package main

import (
	"fmt"
)


func main() {


	// Example 1 - Print 100 numbers
	fmt.Println("Example 1 - Print 100 numbers")
	for i := 0; i <= 100; i ++ {
		fmt.Printf("\"Hello World\" is being printed %d times\n", i);
	}




	// Example 2 - Example of Nested Loop
	fmt.Println("\n\n")
	fmt.Println("Example 2 - Example of Nested Loop")
	for i := 0; i <= 10; i ++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("*")
		}
		fmt.Println("")
	}


	// Example 3 - Example of Nested Loop
	fmt.Println("\n\n")
	fmt.Println("Example 3 - Example of Nested Loop")
	for i := 0; i <= 10; i ++ {
		fmt.Printf("The Outer Loop is %d:\n", i);
		for j := 0; j < 3; j ++ {
			fmt.Printf("\t\tThe Inner Loop is %d\n", j);
		}
	}



	// Example 4 - This is similar to a while loop
	fmt.Println("\n\n")
	fmt.Println("Example 4 - This is similar to a while loop")
	x := 1;
	for x < 10 {
		fmt.Println(x);
		x ++;
	}
	fmt.Println("Done.");



	// Example 5 - for with and if example
	fmt.Println("\n\n")
	fmt.Println("Example 5 - for with and if example")
    i := 0;
	for {
        if i >= 10 {
        	break;
        }
        fmt.Println(i);
        i ++;
	}
	fmt.Println("Done.")


	// Example 6 - break and continue in action
	fmt.Println("\n\n")
	fmt.Println("Example 6 - break and continue in action - Print Even numbers")
	number := 0;
	for {
		number ++;
		if number > 100 {
			break;
		}

		if number % 2 != 0 {
			continue;
		}
		fmt.Printf("%d is a even number\n", number);
	}

	// Example 7 - Print ASCII characters for the first 200 numbers
	fmt.Println("\n\n")
	fmt.Println("Example 7 - Print ASCII characters for the first 200 numbers")
	for character := 33; character <= 122; character ++ {
	    fmt.Printf("The ASCII character of %d is %c\n", character, character);
	}


	/* Example 8 - Simple If statement
       Default condition is always true

       true = true
       false = false
       !true = false
       !false = true
       !!true = true
       !!false = false

       Comparision Operators:

       >; <; <=; >=; ==; !=
	*/ 
	fmt.Println("\n\n")
	fmt.Println("Example 8 - Simple If statement")
	if true {
		fmt.Println("001");
	} else {
		fmt.Println("002");
	}


	if !true {
		fmt.Println("003");
	} else {
		fmt.Println("004");
	}


	if !!true {
		fmt.Println("005");
	} else {
		fmt.Println("006");
	}

	if 2 == 2 {
		fmt.Println("007");
	} else {
		fmt.Println("008");
	}


	if 2 != 2 {
		fmt.Println("009");
	} else {
		fmt.Println("010");
	}


	/* Example 9 - Print two statements in one line

        Ideal case - Compiler puts ";" after every line, but if you need to print two print statements in one line, then it needs to be explicitly separated by a ";"
        In the below example: fmt.Println("Do Something"); fmt.Println("Do Another thing"); - is being separated by a ";" - so it like two statements
	*/
	fmt.Println("\n\n")
	fmt.Println("Example 9 - Print two statements in one line")
	if x := 42; x == 42 {
		fmt.Println(x);
	}

	fmt.Println("Do Something"); fmt.Println("Do Another thing");


	// Example 10 - Example of If-else-If
	fmt.Println("\n\n")
	fmt.Println("Example 10 - Example of If-else-If");
	var integer int;
	fmt.Println("Enter the integer: ");
	fmt.Scanf("%d", &integer);

	if integer % 4 == 0 && integer % 6 == 0 {
		fmt.Println("Linkedin");
	} else if integer % 4 == 0 {
		fmt.Println("Linked");
	} else if integer % 6 == 0 {
		fmt.Println("In");
	} else {
		fmt.Println("Not divisible by either 4 or 6");
	}


	// Example 11 - Example of loops,conditions and modulo put together
	fmt.Println("\n\n")
	fmt.Println("Example 11 - Example of loops,conditions and modulo put together")
	for number := 0; number <= integer; number ++ {
		if number % 4 == 0 && number % 6 == 0 {
			fmt.Printf("The number %d is divisible by both 4 and 6 - LINKEDIN\n", number);
		} else if number % 4 == 0 {
			fmt.Printf("The number %d is divisible by only 4       - LINKED\n", number);
		} else if number % 6 == 0 {
			fmt.Printf("The number %d is divisible by only 6       - IN\n", number);
		} else {
			fmt.Printf("The number %d is not divisible by 4 and 6  - INVALID\n", number);
		}
	}

	/* Example 12 - Switch Examples
	switch has 3 options:
		case
		break
		fallthrough

		fallthrough - Ideally, if a switch condition is satisfied, then it prints / executes functions in that block, and exiists, even if another case statement is true. Fallthrough - even after a case statement is true and executed, it moves control to the next step.

	*/
	fmt.Println("\n\n")
	fmt.Println("Example 12 - Example of Switch Example")

	for {

	var input string;
	fmt.Println(`Select one of the below Tasks: 
			e - Eating
			d - drinking
			q - quit
		`)

    fmt.Scanf("%s", &input);

		if input == `q` {
			fmt.Println("You have exited the applicatiion successfully");
			break;
		}

		switch {
			case input == `e`:
				fmt.Println("You have selected the option 'eating'");
			case input == `d`:
				fmt.Println("You have selected the option 'drinking'");
			default:
				fmt.Println("Your selected option is ", input);
				fmt.Println("your're option is invalid");
			}
		continue;
	}


	/* Example 13 - Example of Fallthrough */
	fmt.Println("\n\n")
	fmt.Println("Example 13 - Example of Fallthrough")

	switch {
		case 1 == 1:
			fmt.Println("This case is true");
			fallthrough;

		case 1 == 2:
			fmt.Println("This case is false, and since follthrough is enabled, this case (even if it is false), will be printed..");

		case 3 == 3:
			fmt.Println("This case is true, and since 1 == 1 is already valid, control wouldnt come here.");
	}


	/* Example 14 - Truth Table */
	fmt.Println("\n\n")
	fmt.Println("Example 14 - Truth Table")


	fmt.Println("true and true: ", true && true);
	fmt.Println("true and false: ", true && false);
	fmt.Println("false and false: ", false && false);


	fmt.Println("true or true: ", true || true);
	fmt.Println("true or false: ", true || false);
	fmt.Println("false or false: ", false || false);


	fmt.Println("!true: ", !true);
	fmt.Println("!false: ", !false);

}

