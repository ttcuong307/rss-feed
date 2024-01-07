-- name: GetFeedFollowsForUser :many
SELECT * FROM feed_follows WHERE user_id = ?;
--

-- name: CreateFeedFollow :exec
INSERT INTO feed_follows (
  id,
  feed_id,
  user_id,
  created_at,
  updated_at
) VALUES (
 ?, ?, ?, ?, ?
);
--

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows WHERE id = ? && user_id = ?;