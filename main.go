package main

import (
	"log"
	"net/http"
)

func main() {
	playerServer := &PlayerServer{
		store: NewInMememoryPlayerStore(),
	}
	handler := http.Handler(playerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
