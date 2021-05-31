package scraper

import twitterscraper "github.com/n0madic/twitter-scraper"

func GetTwitterProfile(username string) *twitterscraper.Profile {
	profile, err := twitterscraper.GetProfile(username)
	if err != nil {
		panic(err)
	}
	return &profile
}
