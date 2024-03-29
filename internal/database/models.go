// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0

package database

import (
	"database/sql"
	"time"
)

type Feed struct {
	ID            string
	Name          string
	Url           string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        string
	LastFetchedAt sql.NullTime
}

type FeedFollow struct {
	ID        string
	FeedID    string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	ID          string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Url         string
	Description sql.NullString
	PublishedAt sql.NullTime
	FeedID      string
}

type User struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ApiKey    string
}
