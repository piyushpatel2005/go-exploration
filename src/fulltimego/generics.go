package main

import "fmt"

type CustomMap[K comparable, V any] struct {
	data map[K]V
}

func (m *CustomMap[K, V]) Insert(k K, v V) error {
	m.data[k] = v
	return nil
}

func NewCustomMap[K comparable, V any]() *CustomMap[K, V] {
	return &CustomMap[K, V]{
		data: make(map[K]V),
	}
}

func foo[T any](value T) {
	fmt.Println(value)
}

func main() {
	m1 := NewCustomMap[string, int]()
	m1.Insert("one", 1)
	m1.Insert("two", 2)

	m2 := NewCustomMap[int, float64]()
	m2.Insert(1, 1.1)
	m2.Insert(2, 2.2)

	foo[int](1)
	foo[string]("hello")
}
