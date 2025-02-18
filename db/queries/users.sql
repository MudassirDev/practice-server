-- name: CreateUser :one
INSERT INTO users (first_name, last_name, created_at, updated_at, email)
VALUES (
    $1,
    $2,
    NOW(),
    NOW(),
    $3
)
RETURNING *;
