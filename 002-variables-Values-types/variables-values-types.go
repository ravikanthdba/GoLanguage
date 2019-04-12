/* short declaration opoerator is used for assigning a variable with a value.
Represented by ":="
Instead of declaration one one line, and assigning on another, this puts together both
Can be assigned with an expression also
*/

/* Best Practise
   - If you want to declare a variable in the function body, then use short declaration.
   - If you want to declare a variable outside the function body, then use "var"

*/


/* 

Data Types:

    - Primitive Data Types - Numerics, Strings, Bools
    - Composite Data Types - Array, Slice, Struct, Pointer, Function, Interfaces, Channel, Map

*/

package main

import "fmt"

var y = 149 // x is declared ouotside the function body, and hence var is used.

var z int // Declare a variable z of the type Int, and assigns a value of 0 (of type INT) to z
var x32 float32
var x64 float64

var string1 = "Hey!!! How are you?"
var string2 = `Hey!!! How are you?`

var string1_anothermethod string = "Hey!!! How are you?"
var string2_anothermethod string = `Hey!!! How are you?`

/*
VARIABLE TYPE  - DEFAULT VALUES

int            - 0
booleans       - False
floats         - 0.0
strings        - ""
pointers       - nil
functions       - nil
interfaces       - nil
slices       - nil
channels       - nil
maps       - nil


If we declare a variable of a specific type, and don't assign any values explicitly, the above values are assigned automatically.
These are called Zero Values

*/


/* 	
    The type of the variable is found out by using the following:
    fmt.Println("%T\n", variable)

    We have declared a variable as Int, we cannot re-assign that variable as String.

    GoLang is a Static Programming Language, and not a dynamic programming language

    String can be declared between:
    	- double quotes
    	- back tick

*/


/*
    We can create our our own type. Example is as follows
*/


type hotdog int;
var mytype hotdog;

func main() {

	a := 100 // a is assiigned a value 100
	fmt.Println("The value of a is", a)

	a = 200 // Re-assiigning a
	fmt.Println("The value of a is", a)

	b := 100 * 2 // assigning with an expression
	fmt.Println("The value of b is", b)

	fmt.Println("The value of y is ", y)

	foo()
	print_z()
	print_floats()
	print_type_of_variable()



	/*
    You can convert the type of the variable from one type to another. In the below example, we will be converting a variablee of type hotdog to integer.
	*/

	type_conversion()


}

func foo() {
	fmt.Println("The value of y from foo is :", y);
}

func print_z() {
	fmt.Println("The value of z is: ", z);
}


func print_floats() {
	fmt.Println("The values of x32 and x64 is: ", x64, x32);
}


func print_type_of_variable() {
	mytype = 10;
	fmt.Printf("The type of variable y is %T\n", y);
	fmt.Printf("The type of variable string1 is %T\n", string1);
	fmt.Printf("The type of variable string2 is %T\n", string2);
	fmt.Printf("The type of variable string1_anothermethod is %T\n", string1_anothermethod);
	fmt.Printf("The type of variable string2_anothermethod is %T\n", string2_anothermethod);
	fmt.Printf("The type of variable string1 is %T\n", mytype);
	fmt.Printf("The value of mytype is %d\n", mytype);
}

func type_conversion() {
	var variable1 int = 100;
	type hotdog int;
	var variable2 hotdog = 200;
	fmt.Printf("The value of variable1 is %d and its type is %T\n", variable1, variable1);
	fmt.Printf("The value of variable2 is %d and its type is %T\n", variable2, variable2);
    fmt.Printf("Lets type convert variable2 and assign it to variable1...");
    fmt.Printf("The value of variable2 is %d and its type is %T\n", int(variable2), int(variable2));
    variable1 = int(variable2)
    fmt.Println("After Conversioni");
    fmt.Printf("The value of variable1 is %d and its type is %T\n", int(variable1), int(variable1));
    fmt.Printf("The value of variable2 is %d and its type is %T\n", int(variable2), int(variable2));

}