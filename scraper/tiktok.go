package scraper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"test.com/test/model"
)

type TiktokResponse struct {
	Stats struct {
		FollowersCount int `json:"followerCount"`
		PostsCount     int `json:"videoCount"`
	} `json:"stats"`
}

func GetTiktokProfile(username string) *model.Profile {
	nodeUrl := os.Getenv("NODE_URL") + "/" + username
	resp, err := http.Get(nodeUrl)
	if err != nil {
		fmt.Println("Error getting tiktok user: ")
		fmt.Println(err)
	}
	defer resp.Body.Close()
	var respBody TiktokResponse
	_ = json.NewDecoder(resp.Body).Decode(&respBody)
	fmt.Println(respBody.Stats.FollowersCount)
	profile := &model.Profile{Username: username,
		FollowersCount: respBody.Stats.FollowersCount, PostsCount: respBody.Stats.PostsCount}
	return profile
}
