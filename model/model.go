package model

import "time"

type Profile struct {
	Username       string    `json:"username"`
	FollowersCount int       `json:"followersCount"`
	PostsCount     int       `json:"postsCount"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
