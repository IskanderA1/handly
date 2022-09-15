-- name: CreateUser :one
INSERT INTO users(
    project_account_id,
    name, 
    uuid
) VALUES (
    $1, $2, $3
) 
RETURNING *;

-- name: GetUserByUUID :one
SELECT * FROM users
WHERE uuid = $1 LIMIT 1;

-- name: GetUserByProjectAccountId :one
SELECT * FROM users
WHERE project_account_id = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET project_account_id = $2,
    name = $3,
    uuid = $4,
    last_update_at = now()
WHERE uuid = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;