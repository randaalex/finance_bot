-- name: CreateProcessedTransaction :one
INSERT INTO processed_transactions (
  hash,
  firefly_id
) VALUES (
  $1, $2
) RETURNING *;

-- name: GetProcessedTransactionByHash :one
SELECT * FROM processed_transactions
WHERE hash = $1
LIMIT 1;

-- name: GetProcessedTransactionByFireflyID :one
SELECT * FROM processed_transactions
WHERE firefly_id = $1
LIMIT 1;
