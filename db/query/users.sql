-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password_hash
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE user_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY user_id
LIMIT $1
OFFSET $2;

-- name: UpdateUsername :one
UPDATE users
SET username = $2
WHERE user_id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users 
WHERE user_id = $1;

