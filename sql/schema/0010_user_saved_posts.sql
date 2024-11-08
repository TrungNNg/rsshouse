-- +goose Up
CREATE TABLE user_saved_posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    saved_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(id),
    saved_post_id UUID NOT NULL REFERENCES saved_posts(id),
    UNIQUE (user_id, saved_post_id)
);

-- +goose Down
DROP TABLE user_saved_posts;