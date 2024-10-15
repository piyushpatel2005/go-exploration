package main

import "fmt"

type WeaponType int

func (w WeaponType) String() string {
	switch w {
	case Axe:
		return "Axe"
	case Sword:
		return "Sword"
	case Hammer:
		return "Hammer"
	case Stick:
		return "Stick"
	default:
		return "Unknown"
	}
}

//const (
//	Axe    WeaponType = 1
//	Sword  WeaponType = 2
//	Hammer WeaponType = 3
//	Stick  WeaponType = 4
//)

const (
	Axe WeaponType = iota
	Sword
	Hammer
	Stick
)

func getDamage(weaponType string) int {
	switch weaponType {
	case "sword":
		return 10
	case "axe":
		return 15
	case "hammer":
		return 20
	default:
		panic("weapon does not exist")
	}
}

func getDamage2(weaponType WeaponType) int {
	switch weaponType {
	case Sword:
		return 10
	case Axe:
		return 15
	case Hammer:
		return 20
	default:
		panic("weapon does not exist")
	}
}

func main() {
	fmt.Println("Damage: ", getDamage("sword"))
	//fmt.Println("Damage: ", getDamage("knife"))
	fmt.Println("Damage: ", getDamage2(Sword))

	fmt.Printf("Damage of %s : %d", Sword, getDamage2(Sword))
}
