package main

import (
	"github.com/Color-Kat/learn-go/backend-rss/internal/database"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	APIKey    string    `json:"api_key"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		Name:      dbUser.Name,
		APIKey:    dbUser.ApiKey,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}
}

type Feed struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	URL       string    `json:"url"`
	UserId    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseFeedToFeed(dbFeed database.Feed) Feed {
	return Feed{
		ID:        dbFeed.ID,
		Name:      dbFeed.Name,
		URL:       dbFeed.Url,
		UserId:    dbFeed.UserID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
	}
}

func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, databaseFeedToFeed(dbFeed))
	}

	return feeds
}
