/*
   Data Structures in GoLang are as follows:

   - Arrays
   - Slices
   - Maps

*/

package main

import (
	"fmt"
)

func main() {

	// Exercise 1 : Create an array
	fmt.Println("Exercise 1 : Create an array")

	var array [10]int
	fmt.Println(array)
	array[4] = 30
	fmt.Println(array)

	fmt.Printf("The length of the array x is %d\n", len(array))

	for i := 0; i <= len(array)-1; i++ {
		fmt.Printf("At index %d in the array 'array', the value is %d\n", i, array[i])
	}

	fmt.Printf("The value at index 3 in the array x is : %d\n", array[2])

	// Exercise 2 : Create an slice
	// a SLICE allows to group together the values of the same type - // COMPOSITE LITERAL

	fmt.Println("\n\n")
	fmt.Println("Exercise 2 : Create an slice")
	slice := []int{4, 5, 6, 7, 8, 42}
	fmt.Println(slice)

	// Exercise 3 : Loop through a slice
	fmt.Println("\n\n")
	fmt.Println("Exercise 3 : Loop through a slice")

	for key, value := range slice {
		fmt.Printf("At index position %d of the slice named 'slice', the value is %d\n", key, value)
	}

	/* We can check the elements of the slice from 2nd position to 4th position. It gives the elements from 02nd position and 03rd position
	   slice[0:2] - elements from 0 and 1 - hence it is always excluding upper bound
	   slice[:2] - means the same as 0:2
	   slice[2:] - means elements from index 2 to last
	   slice[:] - all elements in the list
	   slice[3:len(slice)] - elements from index 3 to last

	   for key, value := range slice[3:len(slice)] - the slice is cut from index 3 to the last element, and looped over. Hence the key will be taken from 0 to 2

	*/
	fmt.Println("The values of the slice from 02nd position to 4th position are:", slice[3:len(slice)])

	for key, value := range slice[3:len(slice)] {
		fmt.Printf("At index %d, the value is %d\n", key, value)
	}

	fmt.Printf("The type of the slice is %T\n", slice)

	// Exercise 4 : Append a value to a slice
	fmt.Println("\n\n")
	fmt.Println("Exercise 4 : Append a value to a slice")

	var number int
	fmt.Println("Enter the element which you need to add to the slice:")
	fmt.Scanf("%d", &number)

	slice = append(slice, number)

	fmt.Printf("The slice after the addition of the number %d is: %d\n", number, slice)

	slice2 := []int{2, 3, 4, 5, 6}
	fmt.Println("The values of slice2 are ", slice2)

	/* We can append two slices here
	   The syntax of appending 2 slices are as follows:

	       slice = append(slice1, slice2...)

	       ... - is called the variadic parameters, which means, you can add unlimited variables..s
	*/

	slice = append(slice, slice2...)
	fmt.Println("The values of slice after appending slice2 are ", slice)

	/* Exercise 5 : Deleting a value to a slice
	   There is no optioin like delete from a slice. If you need to eliminate a particular value, you can do this way:
	   We are removing element 42

	   Next we want to remove 7 and 8
	*/
	fmt.Println("\n\n")
	fmt.Println("Exercise 5 : Deleting a value to a slice")
	slice = append(slice[:5], slice[6:]...)
	fmt.Println("Now the slice is : ", slice)
	slice = append(slice[:3], slice[5:]...)
	fmt.Println("Now the slice is : ", slice)

	/* Exercise 6 : Updating a value to a slice */
	fmt.Println("\n\n")
	fmt.Println("Exercise 6 : Updating a value to a slice")
	slice[4] = 0
	fmt.Println("Now the slice is : ", slice)

	/* Exercise 7 : Creating a slice using make
	   make - is a built-in function, and is used for dynamically increasing the size.

	   Advantage of slice vs make

	   When we create a slice using the slice keyword, we do not define the range of values we are going to store.
	   So For Example, if we exceed the limit of the slice, the entire slice is copied to a new memory location, and the existing slice is deleted.
	   This happens everytime, when a new valuee is added, and this takes some proceessing time


	   In case of a slice using make function, we define a rangee of values, which we might place in the slice. In thee below case,  we create a slice with 10 elements.
	   We can maxiimum extend upto 12 elements, If we exceed 12 elements, then all thee elements in thee slice are now copied into another slice with which is created of 24 elements
	   and the existing slice is deleted.

	   len - built in function - which gives the length of the slice - number of elements in the slice
	   cap - max number of elements which can be fit in the slice

	*/
	fmt.Println("\n\n")
	fmt.Println("Exercise 7 : Creating a slice using make")

	makeslice := make([]int, 10, 12)
	fmt.Println(makeslice)
	fmt.Println(len(makeslice))
	fmt.Println(cap(makeslice))
	makeslice[0] = 99
	makeslice[9] = 432
	makeslice = append(makeslice, 433)
	fmt.Println(len(makeslice))
	fmt.Println(cap(makeslice))
	makeslice = append(makeslice, 434)
	fmt.Println(len(makeslice))
	fmt.Println(cap(makeslice))
	makeslice = append(makeslice, 435)
	fmt.Println(len(makeslice))
	fmt.Println(cap(makeslice))
	makeslice = append(makeslice, 436)
	fmt.Println(len(makeslice))
	fmt.Println(cap(makeslice))
	fmt.Println(makeslice)

	/* Exercise 8 : Creating a Multi-dimensional slice. Its like a matrix
	   matrix 1 - Create a 2 X 2 matrix
	   matrix 2 - Create a 2 X 2 matrix

	   sum_matrix - sum of two martrices
	*/

	fmt.Println("\n\n")
	fmt.Println("Exercise 8 : Matrix addition using multi-dimensional slices")
	matrix_1 := [][]int{{1, 2}, {4, 5}}
	matrix_2 := [][]int{{7, 8}, {10, 11}}
	//sum_matrix := make([]int, 10, 20)

	for i := 0; i < len(matrix_1); i++ {
		for j := 0; j < len(matrix_1); j++ {
			fmt.Println(matrix_1[i][j] + matrix_2[i][j])
		}
	}

	/* Exercise 9 : Creating a Multi-dimensional array. Its like a matrix
	   matrix 1 - Create a 2 X 2 matrix
	   matrix 2 - Create a 2 X 2 matrix

	   sum_matrix - sum of two martrices
	*/

	fmt.Println("\n\n")
	fmt.Println("Exercise 9 : Matrix addition using multi-dimensional arrays")
	var m1 = [5][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {1, 2, 3, 4, 5}}
	var m2 = [5][5]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {1, 2, 3, 4, 5}}
	var summatrix [5][5]int

	for i := 0; i < len(m1); i++ {
		for j := 0; j < len(m2); j++ {
			summatrix[i][j] = m1[i][j] + m2[i][j]
		}
	}

	fmt.Println(summatrix)

	/* Exercise 10 : Maps Data Structure
	   Maps are key-value stores
	   Example: {'name': 'phone_number'}
	   Unorder Lists
	*/

	fmt.Println("\n\n")
	fmt.Println("Exercise 10 : Maps data structure")
	occupation := map[string]string{
		"James": "Site Reliability Engineer",
		"Varun": "Staff Site Reliability Engineer",
		"Ravi":  "Site Reliability Engineer",
	}

	fmt.Println("The occupation of James is ", occupation["James"])
	fmt.Println("The occupation of Varun is ", occupation["Varun"])
	/* In this case, Ravi key is not in the map list, and instead of giving key does not exist, it gives a " ".
	   In this case, we usually need to do a check, which checks for the key and returns the value, if the key exists in the map.
	*/
	fmt.Println("The occupation of Ravi is ", occupation["Ravi"])
	v, ok := occupation["Ravi"]
	fmt.Println("The value of v is : ", v)
	fmt.Println("The value of ok is : ", ok)

	/* Printing only if the key exists in map (or) ok is true, is given as follows: */

	if v, ok := occupation["Ravi"]; ok {
		fmt.Println("this would be printed, if the key exists in the map")
		fmt.Println("---------------------------------------------------")
		fmt.Println("The value of v is : ", v)
		fmt.Println("The value of ok is : ", ok)
	}

	/* Adding the Key to the map */
	fmt.Println("New member Raj has been added to the Organization")
	occupation["Raj"] = "Senior Site Reliability Enginieer"
	for key, value := range occupation {
		fmt.Printf("The value for the key %s is %s\n", key, value)
	}

	/* Updating the value to the key */
	fmt.Println("Varun has been promoted")
	occupation["Varun"] = "Senior Staff Site Reliability Engineer"
	for key, value := range occupation {
		fmt.Printf("The occupation of %s is %s\n", key, value)
	}

	/* Deleting the key from the map

	   Syntax - delete(map, key)
	*/
	fmt.Println("Ravi has resigned the Organization")
	delete(occupation, "Ravi")
	for key, value := range occupation {
		fmt.Println(key, "\t", value)
	}

	/* Join two maps */

	fmt.Println("Print Employee name; Occupation and his age")

	age := map[string]int{
		"James": 37,
		"Ravi":  31,
		"Alex":  29,
	}

	for key := range occupation {
		if v, ok := age[key]; ok {
			fmt.Printf("Employee: %s; Occupation: %s; age %d\n", key, occupation[key], v)
		}
	}

}
