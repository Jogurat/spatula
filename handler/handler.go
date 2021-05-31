package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi"
	"test.com/test/model"
	"test.com/test/postgres"
	"test.com/test/scraper"
)

func HandleTwitter(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	username = strings.ToLower(username)
	profile, err := postgres.CheckCache(username, "twitter")
	if err != nil {
		// There is nothing in the cache, grab user info from scraper & store in DB
		fmt.Println("ERROR JE: ", err)
		if err.Error() == "Nothing stored in cache" {
			fmt.Println("ETO ME U ERR")

			newProfile := scraper.GetTwitterProfile(username)
			profileForInsert := &model.Profile{Username: username, FollowersCount: newProfile.FollowersCount, PostsCount: newProfile.TweetsCount}
			postgres.InsertProfile(profileForInsert, "twitter")
			profile = profileForInsert
		}
	}
	profileJson, _ := json.Marshal(*profile)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(profileJson))
}

func HandleTiktok(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	username = strings.ToLower(username)
	profile, err := postgres.CheckCache(username, "tiktok")
	if err != nil {
		// There is nothing in the cache, grab user info from scraper & store in DB
		fmt.Println("ERROR JE: ", err)
		if err.Error() == "Nothing stored in cache" {
			fmt.Println("ETO ME U ERR")

			newProfile := scraper.GetTiktokProfile(username)
			profileForInsert := &model.Profile{Username: username, FollowersCount: newProfile.FollowersCount, PostsCount: newProfile.PostsCount}
			postgres.InsertProfile(profileForInsert, "tiktok")
			profile = profileForInsert
		}
	}
	profileJson, _ := json.Marshal(*profile)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(profileJson))
}
