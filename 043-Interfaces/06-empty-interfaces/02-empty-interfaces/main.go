package main

import "fmt"

type vehicles interface{}

type vehicle struct {
	MaxSpeed int
	Seats    int
	Color    string
}

type car struct {
	vehicle
	Tyres int
	Jet   bool
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Tyres int
}

func main() {
	bmw := car{}
	mercedes := car{}
	maruti := car{}

	boeing757 := plane{}
	boeing767 := plane{}
	boeing787 := plane{}

	sanger := boat{}
	nautique := boat{}
	malibu := boat{}

	rides := []vehicles{bmw, mercedes, maruti, boeing757, boeing767, boeing787, sanger, nautique, malibu}

	for id, ride := range rides {
		fmt.Println(id, ride)
	}
}
