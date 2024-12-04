-- name: Update :exec
UPDATE consumer_loan_limits
SET
    tenor = COALESCE(sqlc.narg(tenor), tenor),
    amount = COALESCE(sqlc.narg(amount), amount),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
    updated_by = COALESCE(sqlc.narg(updated_by), updated_by)
WHERE
    id = sqlc.arg(id);