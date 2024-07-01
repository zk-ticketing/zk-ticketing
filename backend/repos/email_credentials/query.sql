-- name: GetByIdentityCommitment :one
SELECT *
FROM email_credentials
WHERE identity_commitment = $1
LIMIT 1;

-- name: CreateOrUpdateOne :one
INSERT INTO email_credentials (
        id,
        identity_commitment,
        data,
        issued_at,
        expire_at
    )
VALUES ($1, $2, $3, $4, $5) ON CONFLICT (identity_commitment) DO
UPDATE
SET id = $1,
    data = $3,
    issued_at = $4,
    expire_at = $5
RETURNING *;