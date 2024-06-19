-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1;

-- name: CreateUser :one
INSERT INTO users (
        email,
        identity_commitment,
        encrypted_internal_nullifier,
        encrypted_identity_secret,
        created_at
    )
VALUES ($1, NULL, NULL, NULL, NOW())
RETURNING *;