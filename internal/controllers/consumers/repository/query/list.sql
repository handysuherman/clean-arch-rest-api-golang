-- name: List :many
SELECT * FROM consumers WHERE full_name LIKE ? OR legal_name LIKE ?
ORDER BY created_at DESC
LIMIT ?
OFFSET ?;
