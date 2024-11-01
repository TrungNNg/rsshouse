-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username ,hashed_password)
VALUES (
    gen_random_uuid(),
    NOW(),
    NOW(),
    $1,
    $2
)
RETURNING *;

-- name: ResetUsers :exec
DELETE FROM users;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE $1 = username;