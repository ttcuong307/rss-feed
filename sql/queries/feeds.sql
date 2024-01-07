-- name: CreateFeed :exec
INSERT INTO feeds (id, name, url, last_fetched_at, created_at, updated_at, user_id)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetFeeds :many
SELECT * FROM feeds;