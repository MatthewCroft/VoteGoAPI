package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleVoteCard = []byte(`{
	"id": "1",
	"options": [
		"option1",
		"option2"
	],
	"votes": {
		"option1": 0,
		"option2": 0
	}
}`)

func TestCreateVoteCard(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/votecard", bytes.NewBuffer(exampleVoteCard))

	router.ServeHTTP(w, req)

	require.JSONEq(t, string(exampleVoteCard), w.Body.String())
	if w.Code != 201 {
		t.Fatalf("Status should be 201 created, actual code: %v", w.Code)
	}
}

func TestGetVoteCardById(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	r := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/votecard", bytes.NewBuffer(exampleVoteCard))
	// req2, _ := http.NewRequest("PUT", "/votecard/1?option=option1", nil)
	req2, _ := http.NewRequest("GET", "/votecard/1", nil)

	router.ServeHTTP(w, req1)
	require.JSONEq(t, string(exampleVoteCard), w.Body.String())

	router.ServeHTTP(r, req2)
	require.JSONEq(t, string(exampleVoteCard), r.Body.String())
}

func TestGetVoteCardByIdThrowsNotFound(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()

	req1, _ := http.NewRequest("GET", "votecard/3", nil)
	router.ServeHTTP(w, req1)

	if w.Body.String() != "404 page not found" {
		t.Fatalf("message should return page now found")
	}
	if w.Code != 404 {
		t.Fatalf("Vote Cards that do not exist should through Not Found when trying to update votes")
	}
}

func TestUpdateVoteCount(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	r := httptest.NewRecorder()
	req1, _ := http.NewRequest("POST", "/votecard", bytes.NewBuffer(exampleVoteCard))
	req2, _ := http.NewRequest("PUT", "/votecard/1?option=option1", nil)

	router.ServeHTTP(w, req1)
	router.ServeHTTP(r, req2)

	var expectedUpdatedVoteCard = []byte(`{
		"id": "1",
		"options": [
			"option1",
			"option2"
		],
		"votes": {
			"option1": 1,
			"option2": 0
		}
	}`)

	require.JSONEq(t, string(expectedUpdatedVoteCard), r.Body.String())
}

func TestUpdateVoteCountThrowsNotFoundVoteCard(t *testing.T) {

}

func TestUpdateVoteCountThrowsNotFoundOption(t *testing.T) {

}
