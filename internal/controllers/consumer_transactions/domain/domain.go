package domain

import (
	"context"
)

type Usecase interface {
	Create(ctx context.Context, arg *CreateRequestParams, idempotencyKey *string) (int64, error)
	Update(ctx context.Context, id int64, arg *UpdateRequestParams, idempotencyKey *string) (int64, error)
	FindByID(ctx context.Context, id int64) (*ConsumerTransaction, error)
	List(ctx context.Context, arg *FetchParams) (*ConsumerTransactionList, error)
}
