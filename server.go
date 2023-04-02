package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodGet:
		ps.showScore(w, player)
	case http.MethodPost:
		ps.processWin(w, player)
	}
}

func (ps *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := ps.store.GetPlayerScore(player)
	if score == -1 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (ps *PlayerServer) processWin(w http.ResponseWriter, player string) {
	ps.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
