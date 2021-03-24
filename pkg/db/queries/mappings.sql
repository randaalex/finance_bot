-- name: CreateMapping :one
INSERT INTO mappings (
  transaction_hash,
  transaction_text,
  category_id
) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetMappingByTransactionHash :one
SELECT * FROM mappings
WHERE transaction_hash = $1
LIMIT 1;
