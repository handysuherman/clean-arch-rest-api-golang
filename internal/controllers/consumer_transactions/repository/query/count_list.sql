-- name: CountList :one
SELECT COUNT(id) FROM consumer_transactions WHERE consumer_id = ?;
