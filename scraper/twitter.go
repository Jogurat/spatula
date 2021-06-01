package scraper

import (
	"fmt"

	twitterscraper "github.com/n0madic/twitter-scraper"
)

func GetTwitterProfile(username string) *twitterscraper.Profile {
	profile, err := twitterscraper.GetProfile(username)
	if err != nil {
		fmt.Println("Error getting twitter profile: ", err)
		panic(err)
	}
	return &profile
}
