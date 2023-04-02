package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score, ok := s.scores[name]
	if !ok {
		return -1
	}
	return score
}

func TestGetPlayers(t *testing.T) {
	server := &PlayerServer{
		store: &StubPlayerStore{
			scores: map[string]int{
				"Pepper": 20,
				"Floyd":  10,
			},
		},
	}
	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"

		assertResponseBody(t, got, want)
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"

		assertResponseBody(t, got, want)
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("Huy")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := 404

		assertStatus(t, got, want)
	})
}

func newGetScoreRequest(name string) *http.Request {
	res, _ := http.NewRequest(http.MethodGet, "/players/"+name, nil)
	return res
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d, want status %d", got, want)
	}
}
