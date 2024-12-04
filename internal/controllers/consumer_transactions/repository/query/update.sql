-- name: Update :exec
UPDATE consumer_transactions
SET
    admin_fee_amount = COALESCE(sqlc.narg(admin_fee_amount), admin_fee_amount),
    installment_amount = COALESCE(sqlc.narg(installment_amount), installment_amount),
    otr_amount = COALESCE(sqlc.narg(otr_amount), otr_amount),
    interest_rate = COALESCE(sqlc.narg(interest_rate), interest_rate),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
    updated_by = COALESCE(sqlc.narg(updated_by), updated_by)
WHERE
    id = sqlc.arg(id);