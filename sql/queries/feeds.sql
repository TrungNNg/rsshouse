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
    $10,
    $11,
    $12
)
RETURNING *;

-- name: GetFeedByID :one
SELECT * FROM feeds
WHERE id = $1;