package main

import "fmt"

type NumberStorer interface {
	GetAll() ([]int, error)
	Put(int) error
}

type PostgresNumberStore struct {
}

func (p PostgresNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (p PostgresNumberStore) Put(n int) error {
	fmt.Println("storing number in Postgres")
	return nil
}

type MongoDBNumberStore struct {
}

func (m MongoDBNumberStore) GetAll() ([]int, error) {
	return []int{1, 2, 3}, nil
}

func (m MongoDBNumberStore) Put(n int) error {
	fmt.Println("storing number in MongoDB")
	return nil
}

type ApiServer struct {
	numberStore NumberStorer
}

func main() {
	apiServer := ApiServer{
		numberStore: MongoDBNumberStore{},
	}

	if err := apiServer.numberStore.Put(3); err != nil {
		fmt.Println(err)
		panic(err)
	}
	numbers, err := apiServer.numberStore.GetAll()
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println(numbers)
	}
}
