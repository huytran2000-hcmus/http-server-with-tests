package main

import (
	"log"
	"net/http"
)

type InMememoryPlayerStore struct{}

func (i *InMememoryPlayerStore) GetPlayerScore(name string) int {
	return 123
}

func main() {
	playerServer := &PlayerServer{
		store: &InMememoryPlayerStore{},
	}
	handler := http.Handler(playerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
