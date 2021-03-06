package postgres

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"test.com/test/model"
)

var dbPool *pgxpool.Pool

func GetDB() *pgxpool.Pool {
	if dbPool == nil {
		var err error
		dbPool, err = pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
		if err != nil {
			panic(err)
		}
	}
	return dbPool
}

func InsertProfile(profile *model.Profile, socialNetwork string) error {
	dbConn := GetDB()
	sqlString := "INSERT INTO " + socialNetwork + " (username, followers_count, posts_count) VALUES($1, $2, $3)"
	_, err := dbConn.Exec(context.Background(), sqlString, profile.Username,
		profile.FollowersCount, profile.PostsCount)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func UpdateProfile(profile *model.Profile, socialNetwork string) error {
	dbConn := GetDB()
	sqlString := "UPDATE " + socialNetwork + " SET followers_count=$1, posts_count=$2 WHERE username=$3"
	_, err := dbConn.Exec(context.Background(), sqlString,
		profile.FollowersCount, profile.PostsCount, profile.Username)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func CheckCache(username string, socialNetwork string) (*model.Profile, error) {
	dbConn := GetDB()
	var followersCount, postsCount int
	var updatedAt time.Time
	sqlString := "SELECT followers_count, posts_count, updated_at FROM " + socialNetwork + " WHERE username=$1"
	rows, err := dbConn.Query(context.Background(), sqlString, username)
	if err != nil {
		// Handle DB err
		return nil, err
	}
	rowsReturned := 0
	for rows.Next() {
		err = rows.Scan(&followersCount, &postsCount, &updatedAt)
		rowsReturned++
	}
	fmt.Println("Rows returned: ", rowsReturned)
	if rowsReturned == 0 {
		// No items in cache, get from twitter scraper instead
		fmt.Println("Nothing stored in cache")
		return nil, errors.New("Nothing stored in cache")
	}
	// Check if the cache is older than delta minutes
	delta := 10
	validCacheTime := time.Now().Add(time.Duration(-delta) * time.Minute)
	if updatedAt.Before(validCacheTime) {
		// The cache is too stale, we need to update the DB entry
		fmt.Println("Cache is too stale")
		return nil, errors.New("Cache too stale")

	}
	profile := &model.Profile{Username: username, FollowersCount: followersCount, PostsCount: postsCount, UpdatedAt: updatedAt}
	return profile, nil
}
