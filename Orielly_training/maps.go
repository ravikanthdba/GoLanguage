package main

import (
	"fmt"
)

func status (name string) {
	marks := make(map[string]int)
	marks["Alma"] = 10
	marks["Rohit"] = 90
	if marks[name] < 80 {
		fmt.Println(name + " has failed")
	} else {
		fmt.Println(name + " has passed")
	}
}


func statusUsingCommaOk(name string) {
	marks := make(map[string]int)
	marks["Alma"] = 10
	marks["Rohit"] = 90

	var mark int
	var ok bool

	mark, ok = marks[name]
	if !ok {
		fmt.Println("No grade has been recorded for the name: ", name)
	} else if (mark >= 80) {
		fmt.Println(name + " has passed")
	} else {
		fmt.Println(name + " has failed")
	}
}

func main() {
	var maps = make(map[string]int)
	maps["Good"] = 100
	maps["Average"] = 50
	maps["Bad"] = 0
	fmt.Printf("%#v\n", maps)
	fmt.Printf("%#v\n", maps["Good"])
	fmt.Printf("%#v\n", maps["Average"])
	fmt.Printf("%#v\n", maps["Bad"])

	if maps == nil {
		fmt.Println("Maps is now empty")
	} else {
		fmt.Println("Maps is not empty.")
	}

	fmt.Println("\n")

	var newMap = make(map[string]bool)
	fmt.Printf("%#v\n", newMap)
	if (newMap == nil) {
		fmt.Println("Map is empty")
	} else {
		fmt.Println("Map is not empty")	
	}

	fmt.Println("\n")

	var anotherMap map[int]bool
	if anotherMap == nil {
		fmt.Println("AnotherMap doesn't have any value, it's Null")
	} else {
		fmt.Println("AnotherMap have any value, it's not Null")
	}

	fmt.Println("\n")

	// Zero values
	var zeroMapString = make(map[string]int)
	fmt.Println(zeroMapString["nonExistant"])


	var zeroMapInt = make(map[int]int)
	fmt.Println(zeroMapInt[22])

	fmt.Println("\n")
	status("Alma")
	status("Rohit")
	// Even after the Carl key is not available in the map, the status still shows as failing, which is wrong. Alternative for this is the "comma ok" operation
	status("Carl")


	fmt.Println("\n")
	// commaOK operation

	var counters = make(map[string]int)
	counters["a"] = 3
	counters["b"] = 0

	var value int
	var ok bool

	value, ok = counters["a"]
	fmt.Println(value, ok)

	value, ok = counters["b"]
	fmt.Println(value, ok)

	value, ok = counters["c"]
	fmt.Println(value, ok)

	fmt.Println("\n")
	statusUsingCommaOk("Alma")
	statusUsingCommaOk("Rohit")
	statusUsingCommaOk("Carl")


	// Map Literal
	fmt.Println("\n")
	var ranks = make(map[string]int)
	ranks["Gold"] = 1
	ranks["Silver"] = 2
	ranks["Bronze"] = 3

	fmt.Printf("%#v\n", ranks)

	// Map's built in delete function
	fmt.Println("\n")
	delete(ranks, "Bronze")
	fmt.Printf("%#v\n", ranks)
	delete(ranks, "GOLDS")
	fmt.Printf("%#v\n", ranks)

	fmt.Println("\n")
	// Map's delete function using commaOk
	r, ok := ranks["GOLDS"]
	fmt.Println(r, ok)
	fmt.Println(ranks)

	r, ok = ranks["Silver"]
	fmt.Println(r, ok)
	fmt.Println(ranks)

	fmt.Println("\n")
	// For range with maps
	var nameMarks = make(map[string]int)
	nameMarks["Amy"] = 78
	nameMarks["Ben"] = 30
	nameMarks["Jose"] = 80
	for name, marks := range nameMarks {
		fmt.Println(name, marks)
	}
	fmt.Println("\n")
	for name := range nameMarks {
		fmt.Println(name)
	}
	fmt.Println("\n")
	for _, marks := range nameMarks {
		fmt.Println(marks)
	}

	fmt.Println("\n")
	var orderedNameMarks = map[string]int{"Amy": 78, "Ben": 30, "Jose": 80}
	for name, mark := range orderedNameMarks {
		fmt.Println(name, mark)
	}

	for name := range orderedNameMarks {
		fmt.Println(name)
	}
	for _, mark := range orderedNameMarks {
		fmt.Println(mark)
	}

	// Exercise: https://play.golang.org/p/DNUClFfTnhU

	fmt.Println("==========================================")
}