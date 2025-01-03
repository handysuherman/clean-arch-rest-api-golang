// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: count_list.sql

package repository

import (
	"context"
)

const countList = `-- name: CountList :one
SELECT COUNT(id) FROM affiliated_dealers WHERE affiliated_dealer_name LIKE ?
`

func (q *Queries) CountList(ctx context.Context, affiliatedDealerName string) (int64, error) {
	row := q.db.QueryRowContext(ctx, countList, affiliatedDealerName)
	var count int64
	err := row.Scan(&count)
	return count, err
}
