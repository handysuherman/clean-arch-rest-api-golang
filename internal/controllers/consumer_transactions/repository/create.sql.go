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
INSERT INTO consumer_transactions (
    consumer_id,
    contract_number,
    admin_fee_amount,
    installment_amount,
    otr_amount,
    interest_rate,
    transaction_date,
    affiliated_dealer_id,
    created_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateParams struct {
	ConsumerID         int64               `json:"consumer_id"`
	ContractNumber     string              `json:"contract_number"`
	AdminFeeAmount     decimal.NullDecimal `json:"admin_fee_amount"`
	InstallmentAmount  decimal.NullDecimal `json:"installment_amount"`
	OtrAmount          decimal.NullDecimal `json:"otr_amount"`
	InterestRate       decimal.NullDecimal `json:"interest_rate"`
	TransactionDate    string              `json:"transaction_date"`
	AffiliatedDealerID int64               `json:"affiliated_dealer_id"`
	CreatedAt          string              `json:"created_at"`
}

func (q *Queries) Create(ctx context.Context, arg *CreateParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, create,
		arg.ConsumerID,
		arg.ContractNumber,
		arg.AdminFeeAmount,
		arg.InstallmentAmount,
		arg.OtrAmount,
		arg.InterestRate,
		arg.TransactionDate,
		arg.AffiliatedDealerID,
		arg.CreatedAt,
	)
}
