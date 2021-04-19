-- name: CreateMapping :one
INSERT INTO mappings (
  transaction_details,
  category_id
) VALUES (
  $1, $2
)  ON CONFLICT (transaction_details) DO UPDATE SET category_id = $2
RETURNING *;

-- name: GetMappingByTransactionDetails :one
SELECT * FROM mappings
WHERE transaction_details = $1
LIMIT 1;
