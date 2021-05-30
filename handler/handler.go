package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"test.com/test/postgres"
)

func HandleTwitter(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")

	profile, err := postgres.CheckCache(username)
	if err != nil {
		// There is nothing in the cache, grab user info from scraper & store in DB
		// profile, err = postgres.
	}
	profileJson, _ := json.Marshal(*profile)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(profileJson))
}
