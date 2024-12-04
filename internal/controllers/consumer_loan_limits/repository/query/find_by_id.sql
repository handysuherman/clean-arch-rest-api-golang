-- name: FindByID :one
SELECT * FROM consumer_loan_limits WHERE id = ? LIMIT 1;
