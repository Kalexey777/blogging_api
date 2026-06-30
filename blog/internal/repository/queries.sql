-- name: GetPosts :many
SELECT id, title, content, category, tags, createdAt, updatedAt
FROM posts;

-- name: CreatePost :one
INSERT INTO posts (title, content, category, tags, createdAt, updatedAt)
VALUES ($1, $2, $3, $4, NOW(), NOW())
RETURNING id, title, content, category, tags, createdAt, updatedAt;

-- name: UpdatePostByID :one
UPDATE posts
SET title = sqlc.arg(title),
    content = sqlc.arg(content),
    category = sqlc.arg(category),
    tags = sqlc.arg(tags),
    updatedAt = NOW()
WHERE id = sqlc.arg(id)
RETURNING id, title, content, category, tags, createdAt, updatedAt;

-- name: GetPostByID :one
SELECT id, title, content, category, tags, createdAt, updatedAt
FROM posts
WHERE id = sqlc.arg(id);

-- name: DeletePostByID :one
DELETE
FROM posts
WHERE id = sqlc.arg(id)
RETURNING id, title, content, category, tags, createdAt, updatedAt;
