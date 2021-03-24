-- name: CreateTransaction :one
INSERT INTO processed_transactions (
  id
) VALUES (
  $1
) RETURNING *;

-- name: GetProcessedById :one
SELECT * FROM processed_transactions
WHERE id = $1
LIMIT 1;
