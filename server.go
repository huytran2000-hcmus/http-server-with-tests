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
	name := strings.TrimPrefix(r.URL.Path, "/players/")

	fmt.Fprint(w, ps.store.GetPlayerScore(name))
}

func GetPlayerScore(name string) int {
	if name == "Pepper" {
		return 20
	}

	if name == "Floyd" {
		return 10
	}

	return -1
}
