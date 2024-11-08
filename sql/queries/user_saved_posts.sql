-- name: UserSavePost :exec
INSERT INTO user_saved_posts (
    id, 
    created_at, 
    updated_at, 
    saved_at, 
    user_id, 
    saved_post_id
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3,
    $4
);

-- name: UnsavePost :exec
DELETE FROM user_saved_posts
WHERE user_id = $1 AND saved_post_id = $2;

-- name: CountSaved :one
SELECT COUNT(saved_post_id) FROM user_saved_posts
WHERE saved_post_id = $1;