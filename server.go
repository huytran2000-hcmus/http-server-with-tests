package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(string) int
}

type PlayerServer struct {
	store PlayerStore
}

func (ps *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		ps.showScore(w, r)
	}
}

func (ps *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	name := strings.TrimPrefix(r.URL.Path, "/players/")

	score := ps.store.GetPlayerScore(name)
	if score == -1 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, ps.store.GetPlayerScore(name))
}
