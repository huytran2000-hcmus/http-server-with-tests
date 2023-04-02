package main

import "sync"

type InMememoryPlayerStore struct {
	mu    sync.Mutex
	store map[string]int
}

func NewInMememoryPlayerStore() *InMememoryPlayerStore {
	return &InMememoryPlayerStore{
		store: map[string]int{},
	}
}

func (i *InMememoryPlayerStore) GetPlayerScore(player string) int {
	i.mu.Lock()
	defer i.mu.Unlock()
	return i.store[player]
}

func (i *InMememoryPlayerStore) RecordWin(player string) {
	i.mu.Lock()
	defer i.mu.Unlock()
	i.store[player]++
}
