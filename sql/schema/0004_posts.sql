-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    descrip TEXT NOT NULL,
    post_link TEXT NOT NULL UNIQUE,
    published_parsed TIMESTAMP NOT NULL,
    img_url TEXT NOT NULL,
    img_title TEXT NOT NULL,
    guid TEXT NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;