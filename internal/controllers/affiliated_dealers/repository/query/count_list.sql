-- name: CountList :one
SELECT COUNT(id) FROM affiliated_dealers WHERE affiliated_dealer_name LIKE ?;
