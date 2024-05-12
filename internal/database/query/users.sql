-- name: Create :one
INSERT INTO users (
        email
    )
VALUES ($1)
RETURNING *;

-- name: FindMany :many
SELECT * FROM users;