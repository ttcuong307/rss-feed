package main

import (
	"database/sql"
	"rss-feed/internal/database"
	"time"

	"gopkg.in/guregu/null.v3"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	ApiKey    string    `json:"api_key"`
}

func convUserToUser(user database.CreateUserParams) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

func databaseUserToUser(user database.User) User {
	return User{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		ApiKey:    user.ApiKey,
	}
}

type Feed struct {
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	CreatedAt     time.Time  `json:"created_at"`
	UpdatedAt     time.Time  `json:"updated_at"`
	Url           string     `json:"url"`
	UserID        string     `json:"user_id"`
	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

func convFeedToFeed(feed database.CreateFeedParams) Feed {
	return Feed{
		ID:            feed.ID,
		Name:          feed.Name,
		UserID:        feed.UserID,
		Url:           feed.Url,
		LastFetchedAt: SqlNullTimeToTime(feed.LastFetchedAt),
		CreatedAt:     feed.CreatedAt,
		UpdatedAt:     feed.UpdatedAt,
	}
}

type FeedFollow struct {
	ID        string    `json:"id"`
	FeedID    string    `json:"feed_id"`
	UserID    string    `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func convFeedFollowToFeedFollow(feedFollow database.CreateFeedFollowParams) FeedFollow {
	return FeedFollow{
		ID:        feedFollow.ID,
		FeedID:    feedFollow.FeedID,
		UserID:    feedFollow.UserID,
		CreatedAt: feedFollow.CreatedAt,
		UpdatedAt: feedFollow.UpdatedAt,
	}
}

func databaseFeedFollowsToFeedFollows(feedFollows []database.FeedFollow) []FeedFollow {
	var output []FeedFollow
	for _, feedFollow := range feedFollows {
		output = append(output, FeedFollow{
			ID:        feedFollow.ID,
			FeedID:    feedFollow.FeedID,
			UserID:    feedFollow.UserID,
			CreatedAt: feedFollow.CreatedAt,
			UpdatedAt: feedFollow.UpdatedAt,
		})
	}

	return output
}

func NullTimeToTime(nt null.Time) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

func SqlNullTimeToTime(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
