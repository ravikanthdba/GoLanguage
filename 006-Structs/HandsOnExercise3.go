/*

Create a new type: vehicle.
The underlying type is a struct.
The fields:
doors
color
Create two new types: truck & sedan.
The underlying type of each of these new types is a struct.
Embed the “vehicle” type in both truck & sedan.
Give truck the field “fourWheel” which will be set to bool.
Give sedan the field “luxury” which will be set to bool. solution
Using the vehicle, truck, and sedan structs:
using a composite literal, create a value of type truck and assign values to the fields;
using a composite literal, create a value of type sedan and assign values to the fields.
Print out each of these values.
Print out a single field from each of these values.


*/

package main

import (
	"fmt"
)

type vehicle struct {
	door  int
	color string
	name  string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	truck1 := truck{
		vehicle: vehicle{
			door:  2,
			color: "red",
			name:  "Mahindra",
		},
		fourWheel: true,
	}

	car1 := sedan{
		vehicle: vehicle{
			door:  4,
			color: "blue",
			name:  "Nissan",
		},
		luxury: true,
	}

	truck2 := truck{
		vehicle: vehicle{
			door:  3,
			color: "orange",
			name:  "TATA",
		},
		fourWheel: false,
	}

	car2 := sedan{
		vehicle: vehicle{
			door:  2,
			color: "black",
			name:  "Porsche",
		},
		luxury: false,
	}

	mapCar := map[string]sedan{
		car1.name: car1,
		car2.name: car2,
	}

	for _, value := range mapCar {
		fmt.Printf("The car of the brand %s has %d doors and its color is %s\n", value.name, value.door, value.color)
	}

	mapTruck := map[string]truck{
		truck1.name: truck1,
		truck2.name: truck2,
	}

	for _, value := range mapTruck {
		fmt.Printf("The truck of brand %s has %d doors and its color is %s\n", value.name, value.door, value.color)
	}
}
