-- name: AddPost :exec
INSERT INTO posts (
    id, 
    created_at, 
    updated_at, 
    title, 
    descrip, 
    post_link,
    updated_parsed,
    published_parsed, 
    img_url, 
    img_title, 
    guid, 
    feed_id
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
    $12
);

-- name: GetSubcribedPostsOfUser :many
SELECT posts.* FROM posts
JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY posts.published_parsed DESC;

-- name: GetPostsOfFeed :many
SELECT * FROM posts
WHERE feed_id = $1;

-- name: DeletePostByID :exec
DELETE FROM posts
WHERE id = $1;

-- name: UpdatePostByGuid :exec
UPDATE posts
SET 
    updated_at = $1,
    title = $2,
    descrip = $3,
    post_link = $4,
    updated_parsed = $5,
    published_parsed = $6,
    img_url = $7,
    img_title = $8,
    feed_id = $9
WHERE guid = $10;