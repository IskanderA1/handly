-- name: CreateAccount :one
INSERT INTO accounts(
    account_id,
    project_id, 
    name, 
    token
) VALUES (
    $1, $2, $3, $4
) 
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY account_id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET project_id = $2,
    name = $3,
    token = $4,
    last_update_at = now()
WHERE account_id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE account_id = $1;