package domain

import (
	"context"
)

type Usecase interface {
	Create(ctx context.Context, arg *CreateRequestParams) (int, error)
	Update(ctx context.Context, arg *UpdateRequestParams) error
	FindByID(ctx context.Context, id int) (*Consumer, error)
	List(ctx context.Context, arg *FetchParams) (*ConsumerList, error)
}
