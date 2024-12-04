// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: list.sql

package repository

import (
	"context"
)

const list = `-- name: List :many
SELECT id, consumer_id, contract_number, admin_fee_amount, installment_amount, otr_amount, interest_rate, transaction_date, affiliated_dealer_id, created_at, updated_at, updated_by FROM consumer_transactions WHERE consumer_id = ?
ORDER BY created_at DESC
LIMIT ?
OFFSET ?
`

type ListParams struct {
	ConsumerID int64 `json:"consumer_id"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

func (q *Queries) List(ctx context.Context, arg *ListParams) ([]*ConsumerTransaction, error) {
	rows, err := q.db.QueryContext(ctx, list, arg.ConsumerID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ConsumerTransaction{}
	for rows.Next() {
		var i ConsumerTransaction
		if err := rows.Scan(
			&i.ID,
			&i.ConsumerID,
			&i.ContractNumber,
			&i.AdminFeeAmount,
			&i.InstallmentAmount,
			&i.OtrAmount,
			&i.InterestRate,
			&i.TransactionDate,
			&i.AffiliatedDealerID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.UpdatedBy,
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