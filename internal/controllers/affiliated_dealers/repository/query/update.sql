-- name: Update :exec
UPDATE affiliated_dealers
SET
    affiliated_dealer_name = COALESCE(sqlc.narg(affiliated_dealer_name), affiliated_dealer_name),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
    updated_by = COALESCE(sqlc.narg(updated_by), updated_by),
    is_activated = COALESCE(sqlc.narg(is_activated), is_activated),
    is_activated_at = COALESCE(sqlc.narg(is_activated_at), is_activated_at),
    is_activated_updated_at = COALESCE(sqlc.narg(is_activated_updated_at), is_activated_updated_at)
WHERE
    id = sqlc.arg(id);