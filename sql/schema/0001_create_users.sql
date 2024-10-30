-- +goose Up
CREATE TABLE users(
    name TEXT
);

-- +goose Down
DROP TABLE users;