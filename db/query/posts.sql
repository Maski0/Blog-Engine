-- name: CreatePost :one
INSERT INTO posts(
    title,
    content,
    author_id
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: InsertLikePost :one
INSERT INTO likes (
    post_id,
    author_id
) VALUES (
    $1, $2
) RETURNING *;

-- name: CheckIfLikeExists :one
SELECT EXISTS (SELECT 1 FROM "likes" WHERE "post_id" = $1 AND "author_id" = $2);

-- name: GetLikesForPost :one
SELECT COUNT(*) FROM likes 
WHERE post_id = $1;


-- name: CrateComment :one
INSERT INTO comments(
    post_id,
    author_id,
    content
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: CreateCategorie :one
INSERT INTO categories(
    name
) VALUES (
    $1
) RETURNING *;

-- name: CreateTag :one
INSERT INTO tags(
    name
) VALUES (
    $1
) RETURNING *;