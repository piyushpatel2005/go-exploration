package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

type UserProfile struct {
	Id       int
	Comments []string
	Likes    int
	Friends  []int
}

type Response struct {
	data any
	err  error
}

func handleGetUserProfile(id int) (*UserProfile, error) {
	var (
		respch = make(chan Response, 3)
		wg     = &sync.WaitGroup{}
	)

	go getComments(id, respch, wg)
	go getLikes(id, respch, wg)
	go getFriends(id, respch, wg)
	// doing three requests in go routines
	wg.Add(3)
	wg.Wait()
	close(respch)
	userProfile := &UserProfile{}

	//comments, err := getComments(id)
	//if err != nil {
	//	return nil, err
	//}
	//likes, err := getLikes(id)
	//if err != nil {
	//	return nil, err
	//}
	//friends, err := getFriends(id)
	//if err != nil {
	//	return nil, err
	//}

	for resp := range respch {
		if resp.err != nil {
			return nil, resp.err
		}
		switch msg := resp.data.(type) {
		case int:
			userProfile.Likes = msg
		case []string:
			userProfile.Comments = msg
		case []int:
			userProfile.Friends = msg
		}
	}
	//return &UserProfile{
	//	Comments: comments,
	//	Likes:    likes,
	//	Friends:  friends,
	//}, nil

	return userProfile, nil
}

//func getComments(id int) ([]string, error) {
//	time.Sleep(time.Millisecond * 200)
//	comments := []string{"comment1", "comment2", "comment3"}
//	return comments, nil
//}
//
//func getLikes(id int) (int, error) {
//	time.Sleep(time.Millisecond * 200)
//	return 10, nil
//}
//
//func getFriends(id int) ([]int, error) {
//	time.Sleep(time.Millisecond * 300)
//	return []int{0}, nil
//}

func getComments(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	comments := []string{"comment1", "comment2", "comment3"}
	respch <- Response{
		data: comments,
		err:  nil,
	}
	// work is done, decrement the counter
	wg.Done()
}

func getLikes(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 200)
	respch <- Response{
		data: 10,
		err:  nil,
	}
	wg.Done()
}

func getFriends(id int, respch chan Response, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 300)
	friendIds := []int{11, 32, 323, 23}
	respch <- Response{
		data: friendIds,
		err:  nil,
	}
	wg.Done()
}

func main() {
	start := time.Now()
	userProfile, err := handleGetUserProfile(1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", userProfile)
	fmt.Println("Request took", time.Since(start)) // takes more than 600ms for 3 calls without channels
}
