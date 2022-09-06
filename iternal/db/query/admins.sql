-- name: CreateAdmin :one
INSERT INTO admins(
    username,
    password, 
    full_name 
) VALUES (
    $1, $2, $3
) 
RETURNING *;

-- name: GetAdminByUsername :one
SELECT * FROM admins WHERE username = $1 LIMIT 1;

-- name: ListAdmins :many
SELECT * FROM admins
ORDER BY username
LIMIT $1
OFFSET $2;

-- name: UpdateAdmin :one
UPDATE admins
SET password = $2,
    full_name = $3
WHERE username = $1
RETURNING *;

-- name: DeleteAdmin :exec
DELETE FROM admins WHERE username = $1;