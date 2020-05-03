package main

import (
	"fmt"
)

type Gallons float64
type Litres float64

func main() {
	var carFuel Gallons
	var busFuel Litres

	carFuel = Gallons(10.50678)
	busFuel = Litres(240.12345)

	fmt.Println(carFuel)
	fmt.Println(busFuel)

	// Exercise: https://play.golang.org/p/uo1igDqgBv7
}
