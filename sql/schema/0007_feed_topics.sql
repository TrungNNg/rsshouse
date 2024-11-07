-- +goose Up
CREATE TABLE feed_topics (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
    topic_id UUID NOT NULL REFERENCES topics(id) ON DELETE CASCADE,
    UNIQUE (feed_id, topic_id)
);

-- +goose Down
DROP TABLE feed_topics;