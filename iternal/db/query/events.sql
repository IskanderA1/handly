-- name: CreateEvent :one
INSERT INTO events(
    project_id,
    name,
    event_type
) VALUES (
    $1, $2, $3
) 
RETURNING *;

-- name: GetEvent :one
SELECT * FROM events
WHERE id = $1 LIMIT 1;

-- name: GetEventByName :one
SELECT * FROM events
WHERE name = $1 LIMIT 1;


-- name: ListEventsByProjectId :many
SELECT * FROM events
WHERE project_id = $1
ORDER BY id;

-- name: UpdateEvent :one
UPDATE events
SET name = $2,
    event_type = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1;