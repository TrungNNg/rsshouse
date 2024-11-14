-- name: AddFeed :one
INSERT INTO feeds (
    id, 
    created_at, 
    updated_at, 
    title, 
    descrip, 
    link,
    feed_link, 
    updated_parsed,
    published_parsed, 
    lang, 
    img_url, 
    img_title, 
    feed_type, 
    user_id
) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11,
    $12,
    $13,
    $14
)
RETURNING *;

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE id = $1;

-- name: GetFeedsToFetch :many
SELECT * FROM feeds
WHERE $1 - last_fetched_at >= interval '1 hour';

-- name: MarkFeedFetched :one
UPDATE feeds
SET last_fetched_at = $1,
updated_at = $1
WHERE id = $2
RETURNING *;

-- name: UpdateFeedByID :exec
UPDATE feeds
SET 
    updated_at = $2,
    title = $3,
    descrip = $4,
    link = $5,
    feed_link = $6,
    updated_parsed = $7,
    published_parsed = $8,
    lang = $9,
    img_url = $10,
    img_title = $11,
    feed_type = $12,
    last_fetched_at = $13
WHERE id = $1;

-- name: GetFeedsByTopicID :many
SELECT DISTINCT feeds.*, COUNT(feed_follows.user_id) AS follower_count
FROM feeds
JOIN feed_topics ON feeds.id = feed_topics.feed_id
LEFT JOIN feed_follows ON feeds.id = feed_follows.feed_id
WHERE feed_topics.topic_id = ANY($1::UUID[])
  AND (feeds.lang = $2 OR $2 = '')
  AND (feeds.feed_type = $3 OR $3 = '')
GROUP BY feeds.id
ORDER BY follower_count DESC, feeds.created_at DESC
OFFSET $4
LIMIT $5;