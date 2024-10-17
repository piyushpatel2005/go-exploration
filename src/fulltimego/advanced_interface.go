package main

import (
	"fmt"
	"strings"
)

type Putter interface {
	Put(id int, data any) error
}
type Storage interface {
	Putter
	Get(id int) (any, error)
}

type FooStorage struct {
}

func (f *FooStorage) Get(id int) (any, error) {
	return nil, nil
}

func (f *FooStorage) Put(id int, data any) error {
	return nil
}

type BarStorage struct{}

func (b *BarStorage) Get(id int) (any, error) {
	return nil, nil
}

func (b *BarStorage) Put(id int, data any) error {
	return nil
}

type Server struct {
	store Storage
}

func updateValue(id int, data any, p Putter) error {
	return p.Put(id, data)
}

// Functional Go
type TransformFunc func(s string) string

func transformString(s string, fn TransformFunc) string {
	return fn(s)
}

func UpperCase(s string) string {
	return strings.ToUpper(s)
}

func Prepend(s string) string {
	return "Hello " + s
}

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func main() {
	s := &Server{
		store: &FooStorage{},
	}
	s.store.Get(1)
	s.store.Put(1, "data")

	fmt.Println(transformString("Hello", Prepend))
	fmt.Println(transformString("Hello", UpperCase))
	fmt.Println(transformString("Ajay", Prefixer("Hi ")))
	fmt.Println(transformString("Ajay", Prefixer("Hello ")))

}
