-- name: GetEventRegistrations :many
SELECT *
FROM registrations
WHERE event_id = $1;

-- name: GetRegisteredEventsByEmail :many
SELECT *
FROM registrations
WHERE email = $1;

-- name: GetOneByEventIdAndEmail :one
SELECT *
FROM registrations
WHERE event_id = @event_id
    AND email = @email;