-- name: GetByIdentityCommitment :one
SELECT *
FROM email_credentials
WHERE identity_commitment = $1
LIMIT 1;

-- name: CreateOne :one
INSERT INTO email_credentials (
        id,
        identity_commitment,
        data,
        issued_at,
        expire_at
    )
VALUES ($1, $2, $3, $4, $5)
RETURNING *;