
-- name: CreateLog :one
INSERT INTO logs(
    project_id,
    event_id,
    account_id,
    data
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: GetLog :one
SELECT * FROM logs
WHERE id = $1 LIMIT 1;


-- name: ListLogs :many
SELECT * FROM logs
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteLog :exec
DELETE FROM logs WHERE id = $1;