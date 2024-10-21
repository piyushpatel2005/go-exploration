package main

import "sync/atomic"

//type State struct {
//	mu    sync.Mutex
//	count int
//}
//
//func (s *State) setState(i int) {
//	s.mu.Lock()
//	s.count = i
//	s.mu.Unlock()
//}

type State struct {
	count int32
}

func (s *State) setState(i int) {
	atomic.AddInt32(&s.count, int32(1))
}

func main() {
}
