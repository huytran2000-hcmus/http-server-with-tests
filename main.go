package main

import (
	"log"
	"net/http"
)

func main() {
	playerServer := NewPlayerServer(&InMememoryPlayerStore{})
	handler := http.Handler(playerServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
