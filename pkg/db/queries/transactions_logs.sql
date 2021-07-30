-- name: CreateTransactionsLog :one
INSERT INTO transactions_logs (
  description,
  category_id
) VALUES (
  $1, $2
)  ON CONFLICT (description) DO UPDATE SET category_id = $2
RETURNING *;

-- name: GetTransactionsLogByDescription :one
SELECT * FROM transactions_logs
WHERE description = $1
LIMIT 1;
