-- name: CreateFeed :exec
INSERT INTO feeds (id, name, url, last_fetched_at, created_at, updated_at, user_id)
VALUES (?, ?, ?, ?, ?, ?, ?);

-- name: GetFeeds :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetched_at IS NULL DESC, last_fetched_at DESC
LIMIT ?;

-- name: MarkFeedFetched :exec
UPDATE feeds
SET last_fetched_at = ?, updated_at = ?
WHERE id = ?