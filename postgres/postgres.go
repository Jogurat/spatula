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

func CheckCache(username string) (*model.Profile, error) {
	dbConn := GetDB()
	var followersCount, postsCount int
	var updatedAt time.Time
	rows, err := dbConn.Query(context.Background(),
		"SELECT followers_count, posts_count, updated_at FROM twitter WHERE username=$1", username)
	//.Scan(&followersCount, &postsCount)
	if err != nil {
		// Handle DB err
		return nil, err
	}
	rowsReturned := 0
	for rows.Next() {
		err = rows.Scan(&followersCount, &postsCount, &updatedAt)
		rowsReturned++
	}
	if rowsReturned == 0 { /* or updated_at field is too stale */
		// No items in cache, get from twitter scraper instead
		fmt.Println("Nemam u kesu nista brabo")
		// log.Infof(context.Background(), "Nemam u kesu nista brabo")
		return nil, errors.New("Nothing stored in cache")
	}
	profile := &model.Profile{Username: username, FollowersCount: followersCount, PostsCount: postsCount, UpdatedAt: updatedAt}
	return profile, nil
}
