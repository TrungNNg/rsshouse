-- +goose Up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title TEXT NOT NULL,
    descrip TEXT NOT NULL,
    feed_link TEXT NOT NULL UNIQUE,
    updated_parsed TIMESTAMP NOT NULL,
    lang TEXT NOT NULL,
    img_url TEXT NOT NULL,
    img_title TEXT NOT NULL,
    feed_type TEXT NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id)
);

-- +goose Down
DROP TABLE feeds;