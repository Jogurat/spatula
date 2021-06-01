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

func NewTwitterProfile(username string) *model.Profile {
	newProfile := scraper.GetTwitterProfile(username)
	returnProfile := &model.Profile{Username: username, FollowersCount: newProfile.FollowersCount, PostsCount: newProfile.TweetsCount}
	return returnProfile
}

func HandleTwitter(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	username = strings.ToLower(username)
	profile, err := postgres.CheckCache(username, "twitter")
	if err != nil {
		if err.Error() == "Nothing stored in cache" {
			// There is nothing in the cache, grab user info from scraper & store in DB
			profileForInsert := NewTwitterProfile(username)
			postgres.InsertProfile(profileForInsert, "twitter")
			profile = profileForInsert
		} else if err.Error() == "Cache too stale" {
			// Cache is too stale, update the existing profile by getting the info from the scraper
			profileForUpdate := NewTwitterProfile(username)
			postgres.UpdateProfile(profileForUpdate, "twitter")
			profile = profileForUpdate
		} else {
			// Unknown error
			panic(err)
		}
	}
	profileJson, _ := json.Marshal(*profile)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(profileJson))
}

func NewTiktokProfile(username string) *model.Profile {
	newProfile := scraper.GetTiktokProfile(username)
	returnProfile := &model.Profile{Username: username, FollowersCount: newProfile.FollowersCount, PostsCount: newProfile.PostsCount}
	return returnProfile
}

func HandleTiktok(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "username")
	username = strings.ToLower(username)
	profile, err := postgres.CheckCache(username, "tiktok")
	if err != nil {
		fmt.Println("Error in tiktok handler: ", err)
		if err.Error() == "Nothing stored in cache" {
			// There is nothing in the cache, grab user info from scraper & store in DB

			profileForInsert := NewTiktokProfile(username)
			postgres.InsertProfile(profileForInsert, "tiktok")
			profile = profileForInsert
		} else if err.Error() == "Cache too stale" {
			// Cache is too stale, update the existing profile by getting the info from the scraper
			profileForUpdate := NewTiktokProfile(username)
			postgres.UpdateProfile(profileForUpdate, "tiktok")
			profile = profileForUpdate
		} else {
			// Unknown error
			panic(err)
		}
	}
	profileJson, _ := json.Marshal(*profile)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(profileJson))
}
