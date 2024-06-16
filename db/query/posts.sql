-- name: CreatePost :one
INSERT INTO posts(
    title,
    content,
    author_id
) VALUES (
    $1, $2, $3
) RETURNING *;