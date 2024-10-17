package types

import (
	"github.com/google/uuid"
)

// Person struct represents a person with a name and age
type Address struct {
	Street  string
	City    string
	State   string
	Country string
	Zip     string
}

type Person struct {
	Id      uuid.UUID
	Name    string
	Age     int
	Address Address
}
