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
INSERT INTO consumers (
    nik,
    full_name,
    legal_name,
    birth_place,
    birth_date,
    salary,
    ktp_photo,
    selfie_photo,
    created_at
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
)
`

type CreateParams struct {
	Nik         string              `json:"nik"`
	FullName    string              `json:"full_name"`
	LegalName   sql.NullString      `json:"legal_name"`
	BirthPlace  sql.NullString      `json:"birth_place"`
	BirthDate   sql.NullTime        `json:"birth_date"`
	Salary      decimal.NullDecimal `json:"salary"`
	KtpPhoto    sql.NullString      `json:"ktp_photo"`
	SelfiePhoto sql.NullString      `json:"selfie_photo"`
	CreatedAt   string              `json:"created_at"`
}

func (q *Queries) Create(ctx context.Context, arg *CreateParams) (sql.Result, error) {
	return q.db.ExecContext(ctx, create,
		arg.Nik,
		arg.FullName,
		arg.LegalName,
		arg.BirthPlace,
		arg.BirthDate,
		arg.Salary,
		arg.KtpPhoto,
		arg.SelfiePhoto,
		arg.CreatedAt,
	)
}