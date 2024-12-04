// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: create.sql

package repository

import (
	"context"
	"database/sql"

	"github.com/shopspring/decimal"
)

const create = `-- name: Create :execresult
INSERT INTO consumer_loan_limits (
    consumer_id,
    tenor,
    amount,
    created_at
) VALUES (
    ?, ?, ?, ?
)
`

type CreateParams struct {
	ConsumerID int64           `json:"consumer_id"`
	Tenor      int16           `json:"tenor"`
	Amount     decimal.Decimal `json:"amount"`
	CreatedAt  string          `json:"created_at"`
}

func (q *Queries) Create(ctx context.Context, arg *CreateParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, create,
		arg.ConsumerID,
		arg.Tenor,
		arg.Amount,
		arg.CreatedAt,
	)
}