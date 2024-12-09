-- name: AddSavedPost :one
INSERT INTO saved_posts (id, created_at, updated_at, title, post_link) 
VALUES (
    $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetSavedPostByPostLink :one
SELECT * FROM saved_posts
WHERE post_link = $1;

-- name: DeleteSavedPost :exec
DELETE FROM saved_posts
WHERE id = $1;