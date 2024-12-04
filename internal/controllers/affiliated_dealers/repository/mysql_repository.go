// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"context"
	"database/sql"
)

type Querier interface {
	CountList(ctx context.Context, affiliatedDealerName string) (int64, error)
	Create(ctx context.Context, arg *CreateParams) (sql.Result, error)
	FindByID(ctx context.Context, id int64) (*AffiliatedDealer, error)
	List(ctx context.Context, arg *ListParams) ([]*AffiliatedDealer, error)
	Update(ctx context.Context, arg *UpdateParams) error
}

var _ Querier = (*Queries)(nil)