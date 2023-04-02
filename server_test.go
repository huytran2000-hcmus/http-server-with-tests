package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(player string) int {
	score, ok := s.scores[player]
	if !ok {
		return -1
	}
	return score
}

func (s *StubPlayerStore) RecordWin(player string) {
	s.winCalls = append(s.winCalls, player)
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

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		scores: map[string]int{},
	}
	server := &PlayerServer{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}

		got := store.winCalls[0]
		if got != player {
			t.Errorf("didn't store correct player got %q, want %q", got, player)
		}
	})
}

func newGetScoreRequest(player string) *http.Request {
	res, _ := http.NewRequest(http.MethodGet, "/players/"+player, nil)
	return res
}

func newPostWinRequest(player string) *http.Request {
	res, _ := http.NewRequest(http.MethodPost, "/players/"+player, nil)
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
