package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var jsonContentType = "application/json"

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	http.Handler
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	ps := new(PlayerServer)

	ps.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(ps.leagueHandler))

	router.Handle("/players/", http.HandlerFunc(ps.playerHandler))

	ps.Handler = router
	return ps
}

func (ps *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(ps.store.GetLeague())
}

func (ps *PlayerServer) playerHandler(w http.ResponseWriter, r *http.Request) {
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
