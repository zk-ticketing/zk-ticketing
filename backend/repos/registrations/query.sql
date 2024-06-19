-- name: GetEventRegistrations :many
SELECT *
FROM registrations
WHERE event_id = $1;

-- name: GetRegisteredEventsByEmail :many
SELECT *
FROM registrations
WHERE email = $1;