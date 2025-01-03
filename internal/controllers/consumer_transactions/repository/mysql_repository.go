// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package repository

import (
	"context"
	"database/sql"
)

type Querier interface {
	CountList(ctx context.Context, consumerID int64) (int64, error)
	Create(ctx context.Context, arg *CreateParams) (sql.Result, error)
	CreateAffiliatedDealer(ctx context.Context, arg *CreateAffiliatedDealerParams) (sql.Result, error)
	CreateConsumers(ctx context.Context, arg *CreateConsumersParams) (sql.Result, error)
	FindByID(ctx context.Context, id int64) (*ConsumerTransaction, error)
	List(ctx context.Context, arg *ListParams) ([]*ConsumerTransaction, error)
	Update(ctx context.Context, arg *UpdateParams) error
}

var _ Querier = (*Queries)(nil)
