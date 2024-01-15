-- name: CreatePost :exec
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES (?, ?, ?, ?, ?, ?, ?, ?);
--

-- name: GetPostsForUser :many
SELECT posts.* FROM posts
JOIN feed_follows ON feed_follows.feed_id = posts.feed_id
WHERE feed_follows.user_id = ?
ORDER BY posts.published_at DESC
LIMIT ?;
--