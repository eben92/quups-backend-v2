-- name: CreateUser :one
INSERT INTO users (
        email
    )
VALUES ($1)
RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
    WHERE id = $1
    LIMIT 1;


-- name: GetUserByEmail :one
SELECT * FROM users
    WHERE email = $1
    LIMIT 1;

-- name: GetUserByMsisdn :one
SELECT * FROM users
    WHERE msisdn = $1
    LIMIT 1;

-- name: GetUsers :many
SELECT * FROM users;