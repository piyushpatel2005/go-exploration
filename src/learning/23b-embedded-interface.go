package main

import "fmt"

type Walker interface {
	Walk()
}

type Talker interface {
	Talk()
}

type Flyer interface {
	Fly()
}

type Humanoid interface {
	Walker
	Talker
	Study()
}

type Human struct {
	Name string
	Age  int
}

func (h Human) Walk() {
	fmt.Println(h.Name, "is walking")
}

func (h Human) Talk() {
	fmt.Println(h.Name, "is talking")
}

func (h Human) Study() {
	fmt.Println("Teacher is teaching", h.Name)
	fmt.Println(h.Name, "is studying")
}

type Bird struct {
	Name string
}

func (b Bird) Fly() {
	fmt.Println(b.Name, "is flying")
}

func (b Bird) Talk() {
	fmt.Println(b.Name, "is talking")
}

func (b Bird) Walk() {
	fmt.Println(b.Name, "is walking")
}

func teach(h Humanoid) {
	h.Study()
}

func main() {
	h := Human{"John", 25}
	h.Walk()
	h.Talk()
	teach(h)

	fmt.Println("=====================================")

	b := Bird{"Pigeon"}
	b.Fly()
	b.Talk()
	b.Walk()
	teach(b)
}
