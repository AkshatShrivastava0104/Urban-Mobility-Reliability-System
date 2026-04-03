package main

import "sync"

type Store struct {
	mu     sync.RWMutex
	Driver map[string]int
	Rider  map[string]int
}

func (s *Store) MemoryStore(name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.Driver[name]++

}
