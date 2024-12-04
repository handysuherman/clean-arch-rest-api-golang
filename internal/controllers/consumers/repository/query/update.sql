-- name: Update :exec
UPDATE consumers
SET
    full_name = COALESCE(sqlc.narg(full_name), full_name),
    birth_place = COALESCE(sqlc.narg(birth_place), birth_place),
    birth_date = COALESCE(sqlc.narg(birth_date), birth_date),
    salary = COALESCE(sqlc.narg(salary), salary),
    ktp_photo = COALESCE(sqlc.narg(ktp_photo), ktp_photo),
    selfie_photo = COALESCE(sqlc.narg(selfie_photo), selfie_photo),
    updated_at = COALESCE(sqlc.narg(updated_at), updated_at),
    updated_by = COALESCE(sqlc.narg(updated_by), updated_by),
    is_activated = COALESCE(sqlc.narg(is_activated), is_activated),
    is_activated_at = COALESCE(sqlc.narg(is_activated_at), is_activated_at),
    is_activated_updated_at = COALESCE(sqlc.narg(is_activated_updated_at), is_activated_updated_at)
WHERE
    id = sqlc.arg(id);