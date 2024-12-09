-- name: SubcribeFeed :exec
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5);

-- name: UnsubcribeFeed :exec
DELETE FROM feed_follows 
WHERE user_id = $1 AND feed_id = $2;

-- name: GetSubcribedFeed :many
SELECT feed_id FROM feed_follows
WHERE user_id = $1;