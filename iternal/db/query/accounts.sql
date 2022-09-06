-- name: CreateAccount :one
INSERT INTO accounts(
    account_id,
    name, 
    uuid
) VALUES (
    $1, $2, $3
) 
RETURNING *;

-- name: GetAccount :one
SELECT * FROM accounts
WHERE account_id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM accounts
ORDER BY account_id
LIMIT $1
OFFSET $2;

-- name: UpdateAccount :one
UPDATE accounts
SET name = $2,
    uuid = $3,
    last_update_at = now()
WHERE account_id = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accounts WHERE account_id = $1;