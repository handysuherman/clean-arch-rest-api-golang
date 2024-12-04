-- name: List :many
SELECT * FROM consumer_transactions WHERE consumer_id = ?
ORDER BY created_at DESC
LIMIT ?
OFFSET ?;
