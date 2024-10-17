package main

import "fmt"

type Color int

func (c Color) String() string {
	switch c {
	case RED:
		return "Red"
	case BLUE:
		return "Blue"
	case GREEN:
		return "Green"
	case YELLOW:
		return "Yellow"
	case PURPLE:
		return "Purple"
	default:
		panic("Invalid color")
	}
}

const (
	RED Color = iota
	BLUE
	GREEN
	YELLOW
	PURPLE
)

func main() {
	fmt.Println("The color is", RED)
}
