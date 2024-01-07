// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: feeds.sql

package database

import (
	"context"
	"time"
)

const createFeed = `-- name: CreateFeed :exec
INSERT INTO feeds (id, name, url, created_at, updated_at, user_id)
VALUES (?, ?, ?, ?, ?, ?)
`

type CreateFeedParams struct {
	ID        string
	Name      string
	Url       string
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) error {
	_, err := q.db.ExecContext(ctx, createFeed,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
	)
	return err
}

const getFeeds = `-- name: GetFeeds :many
SELECT id, name, url, created_at, updated_at, user_id FROM feeds
`

func (q *Queries) GetFeeds(ctx context.Context) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getFeeds)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Feed
	for rows.Next() {
		var i Feed
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Url,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}