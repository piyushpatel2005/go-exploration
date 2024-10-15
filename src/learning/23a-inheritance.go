package main

import (
	"fmt"
	"strconv"
)

type Vehicle struct {
	Make  string
	Model string
	Year  int
}

func (v Vehicle) String() string {
	return "Vehicle: " + v.Make + " " + v.Model + " " + strconv.Itoa(v.Year)
}

type ElectricVehicle struct {
	Vehicle
	BatteryCapacity int
}

func (ev ElectricVehicle) String() string {
	return ev.Vehicle.String() + " Battery Capacity: " + strconv.Itoa(ev.BatteryCapacity)
}

func main() {
	v := Vehicle{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2020,
	}
	fmt.Println(v)

	ev := ElectricVehicle{
		BatteryCapacity: 75,
		Vehicle: Vehicle{
			Make:  "Tesla",
			Model: "Model 3",
			Year:  2021,
		},
	}

	fmt.Println(ev)
	fmt.Println(ev.Make)

	ev2 := ElectricVehicle{
		BatteryCapacity: 100,
	}
	ev2.Make = "Tesla"
	ev2.Model = "Model Y"
	ev2.Year = 2022

	fmt.Println(ev2)

}
