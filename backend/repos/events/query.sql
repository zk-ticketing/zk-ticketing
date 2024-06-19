-- name: GetEventByID :one
SELECT *
FROM events
WHERE id = $1;

-- name: ListEvents :many
SELECT *
FROM events;