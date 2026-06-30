-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA IF NOT EXISTS blog;

CREATE TABLE IF NOT EXISTS posts
(
    id SERIAL NOT NULL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    category TEXT NOT NULL,
    tags TEXT[] NOT NULL,
    createdAt TIMESTAMPTZ,
    updatedAt TIMESTAMPTZ
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS posts;
DROP SCHEMA IF EXISTS blog;

-- +goose StatementEnd
