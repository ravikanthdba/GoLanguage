package main

import "fmt"

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
}

type car struct {
	vehicle
	Wheels int
	Door   int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func (v vehicle) information() {
	fmt.Printf("Seats: %v, MaxSpeed: %v, Color: %v\n", v.Seats, v.MaxSpeed, v.Color)
}

func main() {
	toyota := car{}
	tacoma := car{}
	bmw := car{}
	cars := []car{toyota, tacoma, bmw}

	boeing757 := plane{}
	boeing767 := plane{}
	boeing787 := plane{}
	planes := []plane{boeing757, boeing767, boeing787}

	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	boats := []boat{sanger, nautique, malibu}

	for _, plane := range planes {
		fmt.Println(plane)
	}

	for _, car := range cars {
		fmt.Println(car)
	}

	for _, boat := range boats {
		fmt.Println(boat)
	}
}
