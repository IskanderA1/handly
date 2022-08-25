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


-- name: ListEvents :many
SELECT * FROM events
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEvent :one
UPDATE events
SET name = $2,
    event_type = $3
WHERE id = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM events WHERE id = $1;