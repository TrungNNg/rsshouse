-- name: AddPost :exec
INSERT INTO posts (
    id, 
    created_at, 
    updated_at, 
    title, 
    descrip, 
    post_link, 
    published_parsed, 
    img_url, 
    img_title, 
    guid, 
    feed_id
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
    $9
);