-- name: AddFeedTopic :exec
INSERT INTO feed_topics (
    id,
    created_at,
    updated_at,
    feed_id,
    topic_id
) VALUES (
    $1,
    NOW(),
    NOW(),
    $2,
    $3
);