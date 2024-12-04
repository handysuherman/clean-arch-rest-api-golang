// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: list.sql

package repository

import (
	"context"
)

const list = `-- name: List :many
SELECT id, affiliated_dealer_name, created_at, updated_at, updated_by, is_activated, is_activated_at, is_activated_updated_at FROM affiliated_dealers WHERE affiliated_dealer_name LIKE ?
ORDER BY created_at DESC
LIMIT ?
OFFSET ?
`

type ListParams struct {
	AffiliatedDealerName string `json:"affiliated_dealer_name"`
	Limit                int32  `json:"limit"`
	Offset               int32  `json:"offset"`
}

func (q *Queries) List(ctx context.Context, arg *ListParams) ([]*AffiliatedDealer, error) {
	rows, err := q.db.QueryContext(ctx, list, arg.AffiliatedDealerName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*AffiliatedDealer{}
	for rows.Next() {
		var i AffiliatedDealer
		if err := rows.Scan(
			&i.ID,
			&i.AffiliatedDealerName,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UpdatedBy,
			&i.IsActivated,
			&i.IsActivatedAt,
			&i.IsActivatedUpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}