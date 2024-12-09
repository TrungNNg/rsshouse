// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: saved_posts.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const addSavedPost = `-- name: AddSavedPost :one
INSERT INTO saved_posts (id, created_at, updated_at, title, post_link) 
VALUES (
    $1, $2, $3, $4, $5
)
RETURNING id, created_at, updated_at, title, post_link
`

type AddSavedPostParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	Title     string
	PostLink  string
}

func (q *Queries) AddSavedPost(ctx context.Context, arg AddSavedPostParams) (SavedPost, error) {
	row := q.db.QueryRowContext(ctx, addSavedPost,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Title,
		arg.PostLink,
	)
	var i SavedPost
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.PostLink,
	)
	return i, err
}

const deleteSavedPost = `-- name: DeleteSavedPost :exec
DELETE FROM saved_posts
WHERE id = $1
`

func (q *Queries) DeleteSavedPost(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteSavedPost, id)
	return err
}

const getSavedPostByPostLink = `-- name: GetSavedPostByPostLink :one
SELECT id, created_at, updated_at, title, post_link FROM saved_posts
WHERE post_link = $1
`

func (q *Queries) GetSavedPostByPostLink(ctx context.Context, postLink string) (SavedPost, error) {
	row := q.db.QueryRowContext(ctx, getSavedPostByPostLink, postLink)
	var i SavedPost
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Title,
		&i.PostLink,
	)
	return i, err
}
