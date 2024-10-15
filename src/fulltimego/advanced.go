package main

import "fmt"

type Position struct {
	x, y int
}

type Entity struct {
	name    string
	id      string
	version string
	Position
}

type SpecialEntity struct {
	Entity
	specialField float64
}

func main() {
	e := &Entity{
		name:    "MyEntity",
		id:      "id 1",
		version: "1.2",
		Position: Position{
			x: 10,
			y: 15,
		},
	}
	fmt.Printf("%+v\n", e)

	se := &SpecialEntity{
		Entity: Entity{
			name:    "MySpecialEntity",
			id:      "id 2",
			version: "1.3",
			Position: Position{
				x: 20,
				y: 25,
			},
		},
		specialField: 3.14,
	}
	fmt.Printf("%+v\n", se)
	fmt.Println(se.version)
}
