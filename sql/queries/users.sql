-- name: CreateUser :one
INSERT INTO users (name)
VALUES ("Hello")
RETURNING *;