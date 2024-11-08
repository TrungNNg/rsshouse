-- +goose Up
CREATE TABLE saved_posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    post_link TEXT NOT NULL UNIQUE
);

-- +goose Down
DROP TABLE saved_posts;