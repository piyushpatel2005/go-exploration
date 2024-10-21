package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

// If third party API call takes longer, we want to be able to exit or cancel the call, this is where context comes in
func thirdPartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 500) // change timeout and see the results. If below 100 it will pass
	return "some response", nil
}

func fetchUserId() (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()

	type result struct {
		userId string
		err    error
	}

	resultch := make(chan result, 1)

	go func() {
		res, err := thirdPartyHTTPCall()
		resultch <- result{
			userId: res,
			err:    err,
		}
	}()

	select {
	// timeout exceeded or context has been manually cancelled
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-resultch:
		return res.userId, res.err

	}
}

func main() {
	start := time.Now()
	//result, err := thirdPartyHTTPCall()
	//if err != nil {
	//	log.Fatal(err)
	//}

	result, err := fetchUserId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("the response:%v took %v\n", result, time.Since(start))
}
