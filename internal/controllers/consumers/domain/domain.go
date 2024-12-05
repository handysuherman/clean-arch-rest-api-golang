package domain

import (
	"context"
)

type Usecase interface {
	Create(ctx context.Context, arg *CreateConsumerRequestParams) (int, error)
	Update(ctx context.Context, id int, arg *UpdateConsumerRequestParams) error
	FindByID(ctx context.Context, id int) (*Consumer, error)
	List(ctx context.Context, arg *FetchParams) (*ConsumerList, error)
}
