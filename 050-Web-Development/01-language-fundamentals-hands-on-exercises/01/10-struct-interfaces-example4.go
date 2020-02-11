package main

import "fmt"

type vehicle struct {
	door  int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func (t truck) transportationDevice() string {
	return "This is a truck for luggage transport"
}

func (s sedan) transportationDevice() string {
	return "This is a sedan car and is for luxury for people transport"
}

type transport interface {
	transportationDevice() string
}

func report(tr transport) {
	switch tr.(type) {
	case truck:
		fmt.Println("Calling from truck")
	case sedan:
		fmt.Println("Calling from sedan")
	}
	fmt.Println(tr.transportationDevice())
}

func main() {
	t1 := truck{
		vehicle: vehicle{
			door:  4,
			color: "gray",
		},
		fourWheel: true,
	}

	s1 := sedan{
		vehicle: vehicle{
			door:  4,
			color: "green",
		},
		luxury: true,
	}

	fmt.Println("Color of truck: ", t1.color)
	fmt.Println("Number of doors: ", t1.door)
	fmt.Println("fourWheel: ", t1.fourWheel)

	fmt.Println("Color of sedan: ", s1.color)
	fmt.Println("Number of doors: ", s1.door)
	fmt.Println("luxury: ", s1.luxury)

	t1.transportationDevice()
	s1.transportationDevice()

	fmt.Println("Using the report function")

	report(t1)
	report(s1)
}
