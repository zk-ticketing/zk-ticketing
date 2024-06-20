-- name: GetAllByEmail :many
SELECT *
FROM ticket_credentials
WHERE email = $1
ORDER BY issued_at DESC;