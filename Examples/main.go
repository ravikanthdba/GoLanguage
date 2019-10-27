package main

import "fmt"

func main() {
	var a = [6]int{2,3,4,5,6,7}
	fmt.Println(a)
	for value := range a {
		fmt.Println(a[value], &a[value])
	}


	var b = [5]rune{'a', 'b', 'c' ,'d'}
	fmt.Println(b)
	for charvalue := range b {
		fmt.Println(b[charvalue], string(b[charvalue]), &b[charvalue])
	}
	fmt.Print("\n")
	fmt.Println("Example 3")
	var c = []int{}
	fmt.Println(c)
	for sliceValue := range c {
		fmt.Println(sliceValue, c[sliceValue], &sliceValue, &c[sliceValue])
	}
	fmt.Println(" After appending values 1, 2, 3, 4")
	c = append(c, 1)
	c = append(c, 2)
	c = append(c, 3)
	c = append(c, 4)
	fmt.Println(c)
	for sliceValue := range c {
		fmt.Println(sliceValue, c[sliceValue], &sliceValue, &c[sliceValue])
	}
	c = append(c, 7)
	c = append(c, 8)
	fmt.Println("After Appending vaues 7 and 8")
	for sliceValue := range c {
		fmt.Println(sliceValue, c[sliceValue], &sliceValue, &c[sliceValue])
	}

	c = append(c, 10)
	c = append(c, 11)
	fmt.Println("After Appending vaues 10 and 11")
	for sliceValue := range c {
		fmt.Println(sliceValue, c[sliceValue], &sliceValue, &c[sliceValue])
	
	}
}
