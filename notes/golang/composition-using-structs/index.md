# Composition from a struct


Go promotes composition over inheritance. Composition is a way to create a new struct that reuses the properties of an existing struct. This is useful when you want to create a new struct that is similar to an existing struct, but with some additional properties. This is a very common practice in object-oriented programming languages where you define separate classes for each functionality and then you can use those classes as a member for another class. In Go, we don't have a concept of a class. However, we can achieve similar functionality using structs.

Let's define our more generic parent struct called `Person`:

```go
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

func main() {
	v := Vehicle{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2020,
	}
	fmt.Println(v)
}
```

Now let's say you want to define struct for an `ElectricVehicle` that inherits from `Vehicle` and has an additional property called `BatteryCapacity`. In this case, you might have to have similar properties as `Vehicle`. Of course, you can add those properties like any other struct as shown below.

```go
type ElectricVehicle struct {
    Make string
	Model string
	Year int
    BatteryCapacity int
}
```

However, what if you want to define several different types of `Vehicle`. In that case, you may end up with same properties written in multiple different types. It becomes even worse when you have to change one of the property names or types. You will have to change it in multiple places. This is where inheritance comes in handy. You can define a parent struct with common properties and then define child structs that inherit from the parent struct. This way, you can reuse the properties of the parent struct in the child structs.

In Go, you can embed a struct within another struct to achieve inheritance. Let's define a parent struct called `Vehicle` and a child struct called `ElectricVehicle` that inherits from `Vehicle`:

```go
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
		Vehicle: Vehicle{
			Make:  "Tesla",
			Model: "Model 3",
			Year:  2021,
		},
		BatteryCapacity: 75,
	}

	fmt.Println(ev)

}
```

Even though the syntax looks as if we might have to access `Vehicle` properties using `ev.Vehicle.Make`, in reality, Go allows us to access `Vehicle` properties directly on `ElectricVehicle` struct. So, you can access `Make` of `ElectricVehicle` using `ev.Make` directly.

Now, when it comes to initializing the `ElectricVehicle`, you can see that we can either embed `Vehicle` struct directly or we can initialize those properties one by one. Both ways are valid and you can choose the one that you prefer.

```go
ev := ElectricVehicle{
    Vehicle: Vehicle{
        Make:  "Tesla",
        Model: "Model 3",
        Year:  2021,
    },
    BatteryCapacity: 75,
}
```

or

```go
ev2 := ElectricVehicle{
    BatteryCapacity: 100,
}
ev2.Make = "Tesla"
ev2.Model = "Model Y"
ev2.Year = 2022
```

The benefit of embedded structs is that you can access the methods defined on parent struct directly on child struct. In this case, we have defined a `String` method on `Vehicle` struct. We can access this method directly on `ElectricVehicle` struct as shown below:

```go
func (ev ElectricVehicle) String() string {
    return ev.Vehicle.String()
}
```

 Even though embedding a struct looks like inheritance, it's actually not. It is composition.