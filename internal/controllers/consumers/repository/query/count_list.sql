-- name: CountList :one
SELECT COUNT(id) FROM consumers WHERE full_name LIKE ? OR legal_name LIKE ?;
