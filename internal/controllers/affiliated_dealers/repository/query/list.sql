-- name: List :many
SELECT * FROM affiliated_dealers WHERE affiliated_dealer_name LIKE ?
ORDER BY created_at DESC
LIMIT ?
OFFSET ?;
