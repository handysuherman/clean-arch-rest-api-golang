-- name: FindByID :one
SELECT * FROM consumer_transactions WHERE id = ? LIMIT 1;
