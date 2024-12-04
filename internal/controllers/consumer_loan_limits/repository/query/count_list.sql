-- name: CountList :one
SELECT COUNT(id) FROM consumer_loan_limits WHERE consumer_id = ?;
