package main

import (
	"fmt"
	"time"
)

type Server struct {
	quitch chan struct{}
	msgch  chan string
}

func newServer() *Server {
	return &Server{
		quitch: make(chan struct{}),
		msgch:  make(chan string, 128),
	}
}

func (s *Server) Start() {
	fmt.Println("Server started")
	s.Loop()
}

func (s *Server) Loop() {
mainloop:
	for {
		select {
		case <-s.quitch:
			// do some work when we need to quit
			fmt.Println("Server stopped")
			break mainloop
		case msg := <-s.msgch:
			// do something when we have a message
			s.handleMessage(msg)
		default:
			// do some other work
		}
	}
	fmt.Println("Server shutting down")
}

func (s *Server) handleMessage(msg string) {
	fmt.Println("Received message: ", msg)
}

func (s *Server) sendMessage(msg string) {
	s.msgch <- msg
}

func (s *Server) quit() {
	//close(s.quitch)
	s.quitch <- struct{}{}
}

func main() {
	server := newServer()
	//go server.Start()
	//for i := 0; i < 10; i++ {
	//	server.sendMessage(fmt.Sprintf("Message %d", i))
	//}
	go func() {
		time.Sleep(5 * time.Second)
		server.quit()
	}()
	server.Start()
	time.Sleep(1 * time.Second)
}
