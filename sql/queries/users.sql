-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, username ,hashed_password)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;

-- name: ResetUsers :exec
DELETE FROM users;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE $1 = username;