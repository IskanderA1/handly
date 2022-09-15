
-- name: CreateLog :one
INSERT INTO logs(
    project_id,
    event_id,
    user_id,
    data
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: ListProjectLog :many
SELECT * FROM logs
WHERE project_id = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: ListUserLog :many
SELECT * FROM logs
WHERE user_id = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;


-- name: DeleteProjectLogs :exec
DELETE FROM logs WHERE project_id = $1;