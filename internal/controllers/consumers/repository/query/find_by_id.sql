-- name: FindByID :one
SELECT * FROM consumers WHERE id = ? LIMIT 1;
