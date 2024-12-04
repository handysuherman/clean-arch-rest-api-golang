-- name: FindByID :one
SELECT * FROM affiliated_dealers WHERE id = ? LIMIT 1;
