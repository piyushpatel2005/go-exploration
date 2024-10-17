package main

import "fmt"

var (
	floatVar  float32 = 0.1
	floatVar2 float64 = 0.1
	name      string  = "John"
	intVar    int     = 10
	intVar2   int64   = 10
	uintVar   uint32  = 10
	uintVar2  uint64  = 10
	unintVar3 uint8   = 255
	byteVar   byte    = 0x11
	runVar    rune    = 'a'
)

type Player struct {
	name        string
	health      int
	attackPower float64
}

func (player Player) getHealth() int {
	return player.health
}

func getHealth(player Player) int {
	return player.health
}

var weapon string = "AXE"

type Weapon string // custom type

func main() {
	player := Player{
		name:        "John",
		health:      100,
		attackPower: 10.5,
	}
	fmt.Printf("Player: %+v\n", player)
	fmt.Println("Player health: ", getHealth(player))

	// map
	users := map[string]int{
		"John": 10,
		"Jane": 20,
	}
	users["captain"] = 30
	fmt.Println("Users: ", users)
	delete(users, "captain")
	fmt.Println("Users: ", users)

	for key := range users {
		fmt.Println(key)
	}

	// slices
	numbers := []int{1, 2, 3}
	otherNumbers := make([]int, 3)
	otherNumbers[0] = 4
	otherNumbers[1] = 5
	otherNumbers[2] = 6
	fmt.Println("Numbers: ", numbers)

	// arrays
	var numbersArray [3]int
	numbersArray[0] = 1
	fmt.Println(numbersArray)

	// custom types

}
