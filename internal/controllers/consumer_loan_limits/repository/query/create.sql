-- name: Create :execresult
INSERT INTO consumer_loan_limits (
    consumer_id,
    tenor,
    amount,
    created_at
) VALUES (
    ?, ?, ?, ?
);

