// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.18.0
// source: feeds.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createFeed = `-- name: CreateFeed :exec
INSERT INTO feeds (id, name, url, last_fetched_at, created_at, updated_at, user_id)
VALUES (?, ?, ?, ?, ?, ?, ?)
`

type CreateFeedParams struct {
	ID            string
	Name          string
	Url           string
	LastFetchedAt sql.NullTime
	CreatedAt     time.Time
	UpdatedAt     time.Time
	UserID        string
}

func (q *Queries) CreateFeed(ctx context.Context, arg CreateFeedParams) error {
	_, err := q.db.ExecContext(ctx, createFeed,
		arg.ID,
		arg.Name,
		arg.Url,
		arg.LastFetchedAt,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.UserID,
	)
	return err
}

const getFeeds = `-- name: GetFeeds :many
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at FROM feeds
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
			&i.LastFetchedAt,
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

const getNextFeedsToFetch = `-- name: GetNextFeedsToFetch :many
SELECT id, name, url, created_at, updated_at, user_id, last_fetched_at FROM feeds
ORDER BY last_fetched_at IS NULL DESC, last_fetched_at DESC
LIMIT ?
`

func (q *Queries) GetNextFeedsToFetch(ctx context.Context, limit int32) ([]Feed, error) {
	rows, err := q.db.QueryContext(ctx, getNextFeedsToFetch, limit)
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
			&i.LastFetchedAt,
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

const markFeedFetched = `-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = ?, updated_at = ?
WHERE id = ?
`

type MarkFeedFetchedParams struct {
	LastFetchedAt sql.NullTime
	UpdatedAt     time.Time
	ID            string
}

func (q *Queries) MarkFeedFetched(ctx context.Context, arg MarkFeedFetchedParams) error {
	_, err := q.db.ExecContext(ctx, markFeedFetched, arg.LastFetchedAt, arg.UpdatedAt, arg.ID)
	return err
}
