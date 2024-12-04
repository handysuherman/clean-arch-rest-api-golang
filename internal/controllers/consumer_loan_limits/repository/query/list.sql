-- name: List :many
SELECT * FROM consumer_loan_limits WHERE consumer_id = ?
ORDER BY created_at DESC
LIMIT ?
OFFSET ?;
