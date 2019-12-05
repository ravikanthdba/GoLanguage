package main

import (
	"fmt"
	"sort"
)

func main() {
	type people []string
	studyGroup := people{"Ravikanth", "Garimella", "Viraj", "Nandan", "Nagabhushanam", "Swarna Latha", "Bhargavi", "Divya"}
	fmt.Println("Unsorted:", studyGroup)
	sort.Strings(studyGroup)
	fmt.Println(studyGroup)

	s := []string{"Ravikanth", "Garimella", "Viraj", "Nandan", "Nagabhushanam", "Swarna Latha", "Bhargavi", "Divya"}
	fmt.Println("Unsorted:", s)
	sort.Strings(s)
	fmt.Println(s)

	n := []int{9, 3, 4, 5, 1, 7, 8, 6}
	fmt.Println("Unsorted:", n)
	sort.Ints(n)
	fmt.Println(n)

	fmt.Println("Reverse:")

	studyGroup = people{"Ravikanth", "Garimella", "Viraj", "Nandan", "Nagabhushanam", "Swarna Latha", "Bhargavi", "Divya"}
	fmt.Println("Unsorted:", studyGroup)
	sort.Sort(sort.Reverse(sort.StringSlice(studyGroup)))
	fmt.Println("Reverse:", studyGroup)

	s = []string{"Ravikanth", "Garimella", "Viraj", "Nandan", "Nagabhushanam", "Swarna Latha", "Bhargavi", "Divya"}
	fmt.Println("Unsorted:", s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println("Reverse:", s)

	n = []int{9, 3, 4, 5, 1, 7, 8, 6}
	fmt.Println("Unsorted:", n)
	sort.Sort(sort.Reverse(sort.IntSlice(n)))
	fmt.Println("Reverse", n)
}
