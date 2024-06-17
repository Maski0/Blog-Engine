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

-- name: ChekIfLikeExists :one
SELECT EXISTS (SELECT 1 FROM "likes" WHERE "post_id" = $1 AND "author_id" = $2);

-- name: GetLikesForPost :one
SELECT COUNT(*) FROM likes 
WHERE post_id = $1;


-- name: CraeteComment :one
INSERT INTO comments(
    post_id,
    author_id,
    content
) VALUES (
    $1, $2, $3
) RETURNING *;

