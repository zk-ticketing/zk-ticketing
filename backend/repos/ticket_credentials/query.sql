-- name: GetByEventIdAndEmail :one
SELECT *
FROM ticket_credentials
WHERE event_id = @event_id
    AND email = @email
LIMIT 1;

-- name: GetAllByEmail :many
SELECT *
FROM ticket_credentials
WHERE email = $1
ORDER BY issued_at DESC;

-- name: CreateOrUpdateOne :one
INSERT INTO ticket_credentials (
        id,
        email,
        event_id,
        data,
        issued_at,
        expire_at
    )
VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (event_id, email) DO
UPDATE
SET data = $4,
    issued_at = $5,
    expire_at = $6
RETURNING *;