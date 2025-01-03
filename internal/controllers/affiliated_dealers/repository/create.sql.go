// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: create.sql

package repository

import (
	"context"
	"database/sql"
)

const create = `-- name: Create :execresult
INSERT INTO affiliated_dealers (
    affiliated_dealer_name,
    created_at
) VALUES (
    ?, ?
)
`

type CreateParams struct {
	AffiliatedDealerName string `json:"affiliated_dealer_name"`
	CreatedAt            string `json:"created_at"`
}

func (q *Queries) Create(ctx context.Context, arg *CreateParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, create, arg.AffiliatedDealerName, arg.CreatedAt)
}
