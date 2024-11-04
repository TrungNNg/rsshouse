-- name: AddFeed :one
INSERT INTO feeds (
    id, 
    created_at, 
    updated_at, 
    title, 
    descrip, 
    feed_link, 
    updated_parsed, 
    lang, 
    img_url, 
    img_title, 
    feed_type, 
    user_id
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
)
RETURNING *;

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE id = $1;